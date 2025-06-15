package qtr_endtoend

import (
	"encoding/csv"
	"github.com/pkg/profile"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/stats/es_memory"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_golog"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/network/nw_ratelimit"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_replay"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
	"github.com/watermint/toolbox/quality/recipe/qtr_timeout"
	"github.com/watermint/toolbox/resources"
	"io"
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

func Resources() (ui app_ui.UI) {
	bundle := resources.NewBundle()
	lg := esl.Default()
	log.SetOutput(lgw_golog.NewLogWrapper(lg))
	app_resource.SetBundle(bundle)

	mc := app_msg_container_impl.NewContainer()
	if qt_secure.IsSecureEndToEndTest() || app_definitions.IsProduction() {
		return app_ui.NewDiscard(mc, lg)
	} else {
		return app_ui.NewConsole(mc, lg, es_stdout.NewTestOut(), es_dialogue.DenyAll())
	}
}

func MustCreateControl() (ctl app_control.Control, jl app_job.Launcher) {
	ui := Resources()
	
	// Create a unique temporary directory for this test instance to avoid conflicts
	tempDir, err := os.MkdirTemp("", "toolbox-test-*")
	if err != nil {
		panic(err)
	}
	
	// Use the temporary directory as the workspace home
	wb, err := app_workspace.NewBundle(tempDir, app_budget.BudgetUnlimited, esl.ConsoleDefaultLevel(), false, false)
	if err != nil {
		panic(err)
	}
	com := app_opt.Default()
	nop := rc_spec.New(&rc_recipe.Nop{})
	jl = app_job_impl.NewLauncher(ui, wb, com, nop)
	ctl, err = jl.Up()
	if err != nil {
		panic(err)
	}
	return ctl, jl
}

func TestWithDbxClient(t *testing.T, twc func(ctx dbx_client.Client)) {
	TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_client_impl.NewMock("mock", ctl)
		twc(ctx)
	})
}

func TestWithReplayDbxContext(t *testing.T, name string, twc func(ctx dbx_client.Client)) {
	TestWithControl(t, func(ctl app_control.Control) {
		rm, err := qt_replay.LoadReplay(name)
		if err != nil {
			t.Error(err)
			return
		}
		ctx := dbx_client_impl.NewSeqReplayMock(name, ctl, rm)
		twc(ctx)
	})
}

func BenchmarkWithControl(b *testing.B, twc func(ctl app_control.Control)) {
	nw_ratelimit.SetTestMode(true)
	ctl, jl := MustCreateControl()
	
	// Register cleanup to remove temporary directory
	b.Cleanup(func() {
		jl.Down(nil, ctl)
		// Clean up the temporary directory
		if ws := ctl.Workspace(); ws != nil {
			os.RemoveAll(ws.Home())
		}
	})

	twc(ctl.WithFeature(ctl.Feature().AsTest(false)))
}

func TestWithControl(t *testing.T, twc func(ctl app_control.Control)) {
	nw_ratelimit.SetTestMode(true)
	ctl, jl := MustCreateControl()
	
	// Register cleanup to remove temporary directory
	t.Cleanup(func() {
		jl.Down(nil, ctl)
		// Clean up the temporary directory
		if ws := ctl.Workspace(); ws != nil {
			os.RemoveAll(ws.Home())
		}
	})

	twc(ctl.WithFeature(ctl.Feature().AsTest(false)))
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
		_, err := qtr_timeout.RunRecipeTestWithTimeout(ctl, re, true, useMock)
		if pr != nil {
			pr.Stop()
		}
		es_memory.DumpMemStats(l)

		if err == nil {
			return
		}

		if rcErr, _ := qt_errors.ErrorsForTest(l, err); rcErr != nil {
			rs := rc_spec.New(re)
			l.Error("Test failure",
				esl.Error(rcErr),
				esl.String("recipe", rs.CliPath()))
			t.Error(rs.CliPath(), ctl.Workspace().Log(), rcErr)
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
