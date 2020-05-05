package qt_recipe

import (
	"encoding/csv"
	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/profile"
	"github.com/tidwall/gjson"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/go/es_resource"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/log/stats/es_memory"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_golog"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const (
	TestTeamFolderName = "watermint-toolbox-test"
)

func NewTestDropboxFolderPath(rel ...string) mo_path.DropboxPath {
	return mo_path.NewDropboxPath("/" + TestTeamFolderName).ChildPath(rel...)
}

func MustMakeTestFolder(ctl app_control.Control, name string, withContent bool) (path string) {
	path, err := qt_file.MakeTestFolder(name, withContent)
	if err != nil {
		ctl.Log().Error("Unable to create test folder", es_log.Error(err))
		app_exit.Abort(app_exit.FailureGeneral)
	}
	return path
}

func NewTestFileSystemFolderPath(c app_control.Control, name string) mo_path2.FileSystemPath {
	return mo_path2.NewFileSystemPath(MustMakeTestFolder(c, name, true))
}

func NewTestExistingFileSystemFolderPath(c app_control.Control, name string) mo_path2.ExistingFileSystemPath {
	return mo_path2.NewExistingFileSystemPath(MustMakeTestFolder(c, name, true))
}

func resBundle() es_resource.Bundle {
	_, err := rice.FindBox("../../../resources/messages")
	if err == nil {
		return es_resource.New(
			rice.MustFindBox("../../../resources/templates"),
			rice.MustFindBox("../../../resources/messages"),
			rice.MustFindBox("../../../resources/web"),
			rice.MustFindBox("../../../resources/keys"),
			rice.MustFindBox("../../../resources/images"),
			rice.MustFindBox("../../../resources/data"),
		)
	} else {
		// In case the test run from the project root
		return es_resource.New(
			rice.MustFindBox("resources/templates"),
			rice.MustFindBox("resources/messages"),
			rice.MustFindBox("resources/web"),
			rice.MustFindBox("resources/keys"),
			rice.MustFindBox("resources/images"),
			rice.MustFindBox("resources/data"),
		)
	}
}

func Resources() (ui app_ui.UI) {
	bundle := resBundle()
	lg := es_log.Default()
	log.SetOutput(lgw_golog.NewLogWrapper(lg))
	app_resource.SetBundle(bundle)

	mc := app_msg_container_impl.NewContainer()
	if qt_secure.IsSecureEndToEndTest() || app.IsProduction() {
		return app_ui.NewDiscard(mc, lg)
	} else {
		return app_ui.NewConsole(mc, lg, es_stdout.NewDefaultOut(true), es_dialogue.DenyAll())
	}
}

func findTestResource() (resource gjson.Result, found bool) {
	l := es_log.Default()
	p, found := os.LookupEnv("TOOLBOX_TESTRESOURCE")
	if !found {
		return gjson.Parse("{}"), false
	}
	l = l.With(es_log.String("path", p))
	b, err := ioutil.ReadFile(p)
	if err != nil {
		l.Debug("unable to read file", es_log.Error(err))
		return gjson.Parse("{}"), false
	}
	if !gjson.ValidBytes(b) {
		l.Debug("invalid file content", es_log.ByteString("resource", b))
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
	ui := Resources()
	wb, err := app_workspace.NewBundle("", app_budget.BudgetUnlimited, es_log.ConsoleDefaultLevel())
	if err != nil {
		t.Error(err)
		return
	}
	com := app_opt.Default()
	nop := rc_spec.New(&rc_recipe.Nop{})
	jl := app_job_impl.NewLauncher(ui, wb, com, nop)
	ctl, err := jl.Up()
	if err != nil {
		t.Error(err)
		return
	}

	twc(ctl.WithFeature(ctl.Feature().AsTest(false)))

	jl.Down(nil, ctl)
}

func ForkWithName(t *testing.T, name string, c app_control.Control, f func(c app_control.Control) error) {
	err := app_workspace.WithFork(c.WorkBundle(), name, func(fwb app_workspace.Bundle) error {
		cf := c.WithBundle(fwb)
		l := cf.Log()
		l.Info("Execute", es_log.String("name", name))
		return f(cf)
	})
	if re, c := qt_errors.ErrorsForTest(c.Log(), err); !c {
		t.Error(re)
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
		var err error
		if useMock {
			err = re.Test(ctl.WithFeature(ctl.Feature().AsTest(true)))
		} else {
			err = re.Test(ctl.WithFeature(ctl.Feature().AsTest(false)))
		}

		if pr != nil {
			pr.Stop()
		}
		es_memory.DumpMemStats(l)

		if err == nil {
			return
		}

		if re, _ := qt_errors.ErrorsForTest(l, err); re != nil {
			t.Error(re)
		}
	})
}

type RowTester func(cols map[string]string) error

func TestRows(ctl app_control.Control, reportName string, tester RowTester) error {
	l := ctl.Log().With(es_log.String("reportName", reportName))
	csvFile := filepath.Join(ctl.Workspace().Report(), reportName+".csv")

	l.Debug("Start loading report", es_log.String("csvFile", csvFile))

	cf, err := os.Open(csvFile)
	if err != nil {
		l.Warn("Unable to open report CSV", es_log.Error(err))
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
			l.Warn("An error occurred during read report file", es_log.Error(err))
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
				l.Warn("Tester returned an error", es_log.Error(err), es_log.Any("cols", colMap))
				return err
			}
		}
	}

	return nil
}
