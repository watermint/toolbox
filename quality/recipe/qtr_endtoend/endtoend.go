package qtr_endtoend

import (
	"encoding/csv"
	"encoding/json"
	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/profile"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/go/es_resource"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
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
	"github.com/watermint/toolbox/infra/network/nw_replay"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
	"github.com/watermint/toolbox/quality/recipe/qtr_timeout"
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
		ctl.Log().Error("Unable to create test folder", esl.Error(err))
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

func findTestFolder() string {
	l := esl.Default()

	root, err := es_project.DetectRepositoryRoot()
	if err != nil {
		l.Error("Test path not found")
		panic(err)
	}
	return filepath.Join(root, "test")
}

func loadReplay(name string) (rr []nw_replay.Response, err error) {
	l := esl.Default().With(esl.String("name", name))
	tp := findTestFolder()
	rp := filepath.Join(tp, "replay", name)

	l.Debug("Loading replay", esl.String("path", rp))
	b, err := ioutil.ReadFile(rp)
	if err != nil {
		l.Debug("Unable to load", esl.Error(err))
		return nil, err
	}

	if err := json.Unmarshal(b, &rr); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return nil, err
	}

	l.Debug("Replay loaded", esl.Int("numRecords", len(rr)))
	return rr, nil
}

func Resources() (ui app_ui.UI) {
	bundle := resBundle()
	lg := esl.Default()
	log.SetOutput(lgw_golog.NewLogWrapper(lg))
	app_resource.SetBundle(bundle)

	mc := app_msg_container_impl.NewContainer()
	if qt_secure.IsSecureEndToEndTest() || app.IsProduction() {
		return app_ui.NewDiscard(mc, lg)
	} else {
		return app_ui.NewConsole(mc, lg, es_stdout.NewDefaultOut(true), es_dialogue.DenyAll())
	}
}

func TestWithDbxContext(t *testing.T, twc func(ctx dbx_context.Context)) {
	TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context_impl.NewMock(ctl)
		twc(ctx)
	})
}

func TestWithReplayDbxContext(t *testing.T, name string, twc func(ctx dbx_context.Context)) {
	TestWithControl(t, func(ctl app_control.Control) {
		rm, err := loadReplay(name)
		if err != nil {
			t.Error(err)
			return
		}
		ctx := dbx_context_impl.NewReplayMock(ctl, rm)
		twc(ctx)
	})
}

func TestWithControl(t *testing.T, twc func(ctl app_control.Control)) {
	nw_ratelimit.SetTestMode(true)
	ui := Resources()
	wb, err := app_workspace.NewBundle("", app_budget.BudgetUnlimited, esl.ConsoleDefaultLevel())
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
		l.Info("Execute", esl.String("name", name))
		return f(cf)
	})
	if re, c := qt_errors.ErrorsForTest(c.Log(), err); !c && re != nil {
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
		err := qtr_timeout.RunRecipeTestWithTimeout(ctl, re, true, useMock)
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
	l := ctl.Log().With(esl.String("reportName", reportName))
	csvFile := filepath.Join(ctl.Workspace().Report(), reportName+".csv")

	l.Debug("Start loading report", esl.String("csvFile", csvFile))

	cf, err := os.Open(csvFile)
	if err != nil {
		l.Warn("Unable to open report CSV", esl.Error(err))
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
			l.Warn("An error occurred during read report file", esl.Error(err))
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
				l.Warn("Tester returned an error", esl.Error(err), esl.Any("cols", colMap))
				return err
			}
		}
	}

	return nil
}
