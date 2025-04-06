package app_job_impl

import (
	"io"
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

func TestLogFileImpl(t *testing.T) {
	qt_file.TestWithTestFolder(t, "launch", false, func(path string) {
		wb, err := app_workspace.NewBundle(path, app_budget.BudgetUnlimited, esl.LevelInfo, false, false)
		if err != nil {
			t.Error(err)
			return
		}
		mc := app_msg_container_impl.NewSingleWithMessagesForTest(map[string]string{})
		ui := app_ui.NewDiscard(mc, wb.Logger().Logger())
		spec := rc_spec.New(&AppJobTestRecipe{})
		launcher := NewLauncher(ui, wb, app_opt.Default(), spec)
		ctl, err := launcher.Up()
		if err != nil {
			t.Error(err)
		}
		ctl.Log().Debug("Hello")
		launcher.Down(nil, ctl)

		his := NewHistorian(wb.Workspace())
		jobs, err := his.Histories()
		if err != nil {
			t.Error(err)
		}
		if len(jobs) != 1 {
			t.Error(jobs)
		}

		job := jobs[0]

		logs, err := job.Logs()
		if err != nil {
			t.Error(err)
		}
		for _, lf := range logs {
			if lf.Name() == "" {
				t.Error(lf.Name())
			}
			if lf.Type() == "" {
				t.Error(lf.Type())
			}
			if lf.Path() == "" {
				t.Error(lf.Path())
			}
			lf.IsCompressed()
			if err := lf.CopyTo(io.Discard); err != nil {
				t.Error(err)
			}
		}
	})
}
