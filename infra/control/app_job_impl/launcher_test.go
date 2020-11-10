package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

type AppJobTestRecipe struct {
}

func (z AppJobTestRecipe) Preset() {
}

func (z AppJobTestRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z AppJobTestRecipe) Test(c app_control.Control) error {
	return nil
}

func TestLaunchImpl_UpDown(t *testing.T) {
	qt_file.TestWithTestFolder(t, "launch", false, func(path string) {
		wb, err := app_workspace.NewBundle(path, app_budget.BudgetUnlimited, esl.LevelInfo, false)
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
		launcher.Down(nil, ctl)
	})
}
