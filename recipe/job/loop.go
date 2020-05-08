package job

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"time"
)

type Loop struct {
	rc_recipe.RemarkExperimental
	rc_recipe.RemarkConsole
	Until                       mo_time.Time
	IntervalSeconds             mo_int.RangeInt
	QuitOnError                 bool
	RunbookPath                 mo_path2.FileSystemPath
	ErrorRunBookNotFound        app_msg.Message
	ErrorInvalidRunBookContent  app_msg.Message
	ErrorRunBookFailure         app_msg.Message
	ProgressWaitingNextInterval app_msg.Message
	ProgressLoopFinished        app_msg.Message
	ProgressTerminateOnError    app_msg.Message
}

func (z *Loop) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	rb, found := app_workflow.NewRunBook(z.RunbookPath.Path())
	if !found {
		ui.Error(z.ErrorRunBookNotFound.With("Path", z.RunbookPath.Path()))
		return errors.New("runbook not found")
	}
	if err := rb.Verify(c); err != nil {
		ui.Error(z.ErrorInvalidRunBookContent.With("Path", z.RunbookPath.Path()).With("Error", err))
		return err
	}

	for {
		is := time.Now()
		ie := is.Add(time.Duration(z.IntervalSeconds.Value()) * 1000 * time.Millisecond)
		if is.After(z.Until.Time()) {
			l.Debug("Finished", esl.String("now", is.String()), esl.String("until", z.Until.Time().String()))
			ui.Info(z.ProgressLoopFinished)
			return nil
		}

		if err := rb.Run(c); err != nil {
			ui.Error(z.ErrorRunBookFailure.With("Error", err))
			if z.QuitOnError {
				ui.Info(z.ProgressTerminateOnError)
				return err
			}
		}

		ui.Info(z.ProgressWaitingNextInterval.With("Next", ie.Format(time.RFC3339)))
		for {
			if ie.Before(time.Now()) {
				l.Debug("Unsuspend from interval time")
				break
			}
			time.Sleep(1 * 1000 * time.Millisecond)
		}
	}
}

func (z *Loop) Test(c app_control.Control) error {
	// Can't test from this func. Test on tbx_test
	return qt_errors.ErrorScenarioTest
}

func (z *Loop) Preset() {
	z.IntervalSeconds.SetRange(1, 86400*365, 180)
}
