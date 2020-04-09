package qt_recipe

import (
	"encoding/csv"
	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/profile"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_memory"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg_impl"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	TestTeamFolderName = "watermint-toolbox-test"
)

func NewTestDropboxFolderPath(rel ...string) mo_path.DropboxPath {
	return mo_path.NewDropboxPath("/" + TestTeamFolderName).ChildPath(rel...)
}

func NewTestFileSystemFolderPath(c app_control.Control, name string) mo_path.FileSystemPath {
	return mo_path.NewFileSystemPath(qt_file.MustMakeTestFolder(c, name, true))
}

func Resources(t *testing.T) (bx, web *rice.Box, mc app_msg_container.Container, ui app_ui.UI) {
	bx = rice.MustFindBox("../../../resources")
	web = rice.MustFindBox("../../../web")

	mc = app_msg_container_impl.NewContainer(bx)
	ui = app_ui.NewNullConsole(mc, qt_missingmsg_impl.NewMessageTest(t))
	return
}

func findTestResource() (resource gjson.Result, found bool) {
	l := app_root.Log()
	p, found := os.LookupEnv("TOOLBOX_TESTRESOURCE")
	if !found {
		return gjson.Parse("{}"), false
	}
	l = l.With(zap.String("path", p))
	b, err := ioutil.ReadFile(p)
	if err != nil {
		l.Debug("unable to read file", zap.Error(err))
		return gjson.Parse("{}"), false
	}
	if !gjson.ValidBytes(b) {
		l.Debug("invalid file content", zap.ByteString("resource", b))
		return gjson.Parse("{}"), false
	}
	return gjson.ParseBytes(b), true
}

func TestWithApiContext(t *testing.T, twc func(ctx dbx_context.Context)) {
	TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context_impl.NewMock(ctl)
		twc(ctx)
	})
}

func TestWithControl(t *testing.T, twc func(ctl app_control.Control)) {
	nw_ratelimit.SetTestMode(true)
	bx, web, mc, ui := Resources(t)

	cat := rc_catalogue.NewCatalogue([]rc_recipe.Recipe{}, []rc_recipe.Recipe{}, []interface{}{}, []app_feature.OptIn{})
	ctl := app_control_impl.NewSingle(ui, bx, web, mc, cat)
	cs := ctl.(*app_control_impl.Single)
	if res, found := findTestResource(); found {
		var err error
		ctl, err = cs.NewTestControl(res)
		if err != nil {
			t.Error("Unable to create new test control", err)
			return
		}
	}
	err := ctl.Up(app_control.Test(), app_control.Concurrency(runtime.NumCPU()))
	if err != nil {
		os.Exit(app_control.FatalStartup)
	}
	defer ctl.Down()

	twc(ctl)
}

// Returns nil even err != nil if the error type is ignorable.
func RecipeError(l *zap.Logger, err error) (resolvedErr error, cont bool) {
	if err == nil {
		return nil, true
	}
	switch err {
	case qt_errors.ErrorNoTestRequired:
		l.Debug("Skip: No test required for this recipe")
		return nil, false

	case qt_errors.ErrorHumanInteractionRequired:
		l.Debug("Skip: Human interaction required for this test")
		return nil, false

	case qt_errors.ErrorNotEnoughResource:
		l.Debug("Skip: Not enough resource")
		return nil, false

	case qt_errors.ErrorScenarioTest:
		l.Debug("Skip: Implemented as scenario test")
		return nil, false

	case qt_errors.ErrorImplementMe:
		l.Debug("Test is not implemented for this recipe")
		return nil, false

	case qt_errors.ErrorMock:
		l.Debug("Mock test")
		return nil, false

	default:
		return err, false
	}
}

func TestRecipe(t *testing.T, re rc_recipe.Recipe) {
	DoTestRecipe(t, re, false)
}

func DoTestRecipe(t *testing.T, re rc_recipe.Recipe, useMock bool) {
	type Stopper interface {
		Stop()
	}
	nw_ratelimit.SetTestMode(true)
	TestWithControl(t, func(ctl app_control.Control) {
		l := ctl.Log()
		l.Debug("Start testing")

		var pr Stopper
		if !testing.Short() {
			pr = profile.Start(
				profile.ProfilePath(ctl.Workspace().Log()),
				profile.MemProfile,
			)
		}
		if useMock {
			if c, ok := ctl.(app_control.ControlTestExtension); ok {
				c.SetTestValue(qt_endtoend.CtlTestExtUseMock, true)
			}
		}

		err := re.Test(ctl)

		if pr != nil {
			pr.Stop()
		}
		ut_memory.DumpStats(l)

		if err == nil {
			return
		}

		if re, _ := RecipeError(l, err); re != nil {
			t.Error(re)
		}
	})
}

type RowTester func(cols map[string]string) error

func TestRows(ctl app_control.Control, reportName string, tester RowTester) error {
	l := ctl.Log().With(zap.String("reportName", reportName))
	csvFile := filepath.Join(ctl.Workspace().Report(), reportName+".csv")

	l.Debug("Start loading report", zap.String("csvFile", csvFile))

	cf, err := os.Open(csvFile)
	if err != nil {
		l.Warn("Unable to open report CSV", zap.Error(err))
		return err
	}
	defer cf.Close()
	csf := csv.NewReader(cf)
	var header []string
	isFirstLine := true

	for {
		cols, err := csf.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Warn("An error occurred during read report file", zap.Error(err))
			return err
		}
		if isFirstLine {
			header = cols
			isFirstLine = false
		} else {
			colMap := make(map[string]string)
			for i, h := range header {
				colMap[h] = cols[i]
			}
			if err := tester(colMap); err != nil {
				l.Warn("Tester returned an error", zap.Error(err), zap.Any("cols", colMap))
				return err
			}
		}
	}

	return nil
}
