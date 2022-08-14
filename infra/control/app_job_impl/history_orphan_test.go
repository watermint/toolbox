package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func TestOrphanHistory_AppName(t *testing.T) {
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

		oh, found := NewOrphanHistory(ctl.Workspace().Log())
		if !found {
			t.Error(oh, found)
		}

		if oh.JobId() != ctl.Workspace().JobId() {
			t.Error(oh.JobId(), ctl.Workspace().JobId())
		}

		if !oh.IsOrphaned() {
			t.Error(oh.IsOrphaned())
		}
		if oh.IsNested() {
			t.Error(oh.IsNested())
		}
		if oh.AppName() == "" {
			t.Error(oh.AppName())
		}
		if oh.AppVersion() == "" {
			t.Error(oh.AppVersion())
		}
		if oh.JobPath() == "" {
			t.Error(oh.JobPath())
		}
		if oh.RecipeName() != "github.com watermint toolbox infra control app_job_impl app_job_test_recipe" {
			t.Error(oh.RecipeName())
		}
		if v, found := oh.TimeStart(); !found {
			t.Error(v, found)
		}
		if v, found := oh.TimeFinish(); !found {
			t.Error(v, found)
		}
	})
}
