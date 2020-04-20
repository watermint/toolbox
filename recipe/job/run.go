package job

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_process"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"time"
)

type Run struct {
	Fork                    bool
	TimeoutSeconds          mo_int.RangeInt
	RunbookPath             mo_path2.FileSystemPath
	ErrorRunBookNotFound    app_msg.Message
	ErrorTimeoutRequireFork app_msg.Message
	ErrorUnableToFork       app_msg.Message
}

func (z *Run) execFork(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	l.Info("Fork", zap.String("cmd", os.Args[0]), zap.String("runbook", z.RunbookPath.Path()))
	cmd := exec.Command(os.Args[0], "job", "run", "-runbook-path", z.RunbookPath.Path())
	pl := ut_process.NewLogger(cmd, c)
	pl.Start()
	defer pl.Close()

	l.Debug("Start")
	err := cmd.Start()
	if err != nil {
		ui.Error(z.ErrorUnableToFork.With("Error", err))
		return err
	}

	if z.TimeoutSeconds.Value() < 1 {
		l.Info("Waiting for finish process")
		if err := cmd.Wait(); err != nil {
			l.Info("The process finished with an error", zap.Error(err))
		} else {
			l.Info("The process finished")
		}
		return err
	}

	running := true
	go func() {
		cmd.Wait()
		l.Debug("The process finished")
		running = false
	}()

	timeout := time.Now().Add(time.Duration(z.TimeoutSeconds.Value()) * 1000 * time.Millisecond)
	l.Info("Waiting for process", zap.String("timeout", timeout.Format(time.RFC3339)))
	for {
		time.Sleep(500 * time.Microsecond)
		if !running {
			return nil
		}
		if time.Now().After(timeout) {
			l.Debug("Execution timeout, try send kill signal to the process")
			err = cmd.Process.Kill()
			l.Debug("Signal sent", zap.Error(err))
			cmd.Process.Release()
			return nil
		}
	}
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
	ui := c.UI()
	if !z.Fork && z.TimeoutSeconds.Value() > 0 {
		ui.Error(z.ErrorTimeoutRequireFork)
		return errors.New("-timeout-seconds option requires fork")
	}

	if z.Fork {
		return z.execFork(c)
	}
	return z.execInProcess(c)
}

func (z *Run) Test(c app_control.Control) error {
	// Can't test from this func. Test on tbx_test
	return qt_errors.ErrorScenarioTest
}

func (z *Run) Preset() {
	z.Fork = false
	z.TimeoutSeconds.SetRange(0, 86400*365, 0)
}
