package job

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Loop struct {
	Until                       mo_time.Time
	IntervalSeconds             int
	QuitOnError                 bool
	RunbookPath                 mo_path.FileSystemPath
	ErrorRunBookNotFound        app_msg.Message
	ErrorInvalidRunBookContent  app_msg.Message
	ErrorInvalidIntervalSeconds app_msg.Message
	ErrorRunBookFailure         app_msg.Message
	ProgressWaitingNextInterval app_msg.Message
	ProgressLoopFinished        app_msg.Message
	ProgressTerminateOnError    app_msg.Message
}

func (z *Loop) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	rb, found := app_workflow.NewRunBook(z.RunbookPath.Path())
	if z.IntervalSeconds < 1 {
		ui.Error(z.ErrorInvalidIntervalSeconds)
		return errors.New("invalid interval seconds")
	}
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
		ie := is.Add(time.Duration(z.IntervalSeconds) * time.Second)
		if is.After(z.Until.Time()) {
			l.Debug("Finished", zap.String("now", is.String()), zap.String("until", z.Until.Time().String()))
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
			time.Sleep(1 * time.Second)
		}
	}
}

func (z *Loop) Test(c app_control.Control) error {
	l := c.Log()
	p := c.Workspace().Test()
	if err := os.MkdirAll(p, 0755); err != nil {
		l.Debug("Unable to create dir", zap.Error(err))
		return err
	}
	rbPath := filepath.Join(p, "loop.runbook")
	rb := &app_workflow.RunBook{
		Version: "1",
		Steps: []*app_workflow.RunStep{
			{
				Name: "echo-hello",
				Args: []string{"dev", "echo", "-text", "Hello"},
			},
		},
	}
	rbContent, err := json.Marshal(rb)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return err
	}

	if err = ioutil.WriteFile(rbPath, rbContent, 0644); err != nil {
		l.Debug("Unable to write", zap.Error(err))
		return err
	}
	defer os.Remove(rbPath)

	return rc_exec.Exec(c, &Loop{}, func(r rc_recipe.Recipe) {
		m := r.(*Loop)
		m.IntervalSeconds = 1
		m.RunbookPath = mo_path.NewFileSystemPath(rbPath)
		mu := m.Until.(*mo_time.TimeImpl)
		if err = mu.UpdateTime(time.Now().Add(2 * time.Second).Format("2006-01-02T15:04:05")); err != nil {
			panic(err)
		}
	})
}

func (z *Loop) Preset() {
	z.IntervalSeconds = 180
}
