package job

import (
	"errors"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Run struct {
	rc_recipe.RemarkExperimental
	rc_recipe.RemarkConsole
	RunbookPath          mo_path2.FileSystemPath
	ErrorRunBookNotFound app_msg.Message
}

func (z *Run) execInProcess(c app_control.Control) error {
	ui := c.UI()
	rb, found := app_workflow.NewRunBook(z.RunbookPath.Path())
	if !found {
		ui.Error(z.ErrorRunBookNotFound.With("Path", z.RunbookPath.Path()))
		return errors.New("runbook not found")
	}
	if err := rb.Verify(c); err != nil {
		c.Log().Debug("Verification failure")
		return err
	}
	return rb.Run(c)
}

func (z *Run) Exec(c app_control.Control) error {
	return z.execInProcess(c)
}

func (z *Run) Test(c app_control.Control) error {
	// Can't test from this func. Test on tbx_test
	return qt_errors.ErrorScenarioTest
}

func (z *Run) Preset() {
}
