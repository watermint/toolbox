package desktop

import (
	"github.com/andybrewer/mack"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_process"
	"go.uber.org/zap"
	"os/exec"
	"runtime"
	"time"
)

// Tell Dropbox to quit, but no guarantee of stop the process.
type Stop struct {
	WaitSeconds int
}

func (z *Stop) stopDarwin(c app_control.Control) error {
	l := c.Log()
	r, err := mack.Tell("Dropbox", "quit")
	if err != nil {
		l.Error("Unable to send quit", zap.Error(err))
		return nil
	}
	l.Info("Quit", zap.String("response", r))
	return nil
}

func (z *Stop) stopWindows(c app_control.Control) error {
	l := c.Log()
	cmd := exec.Command("taskkill", "/im", "Dropbox.exe", "/f")
	pl := ut_process.NewLogger(cmd, c)
	pl.Start()
	defer pl.Close()
	err := cmd.Start()
	if err != nil {
		l.Error("Unable to start `taskkill`", zap.Error(err))
		return nil
	}
	cmd.Wait()
	return nil
}

func (z *Stop) Exec(c app_control.Control) error {
	if z.WaitSeconds > 0 {
		c.Log().Info("Waiting for stop", zap.Int("seconds", z.WaitSeconds))
		time.Sleep(time.Duration(z.WaitSeconds) * time.Second)
	}
	switch runtime.GOOS {
	case "windows":
		return z.stopWindows(c)
	case "darwin":
		return z.stopDarwin(c)
	default:
		c.Log().Info("Skip")
		return nil
	}
}

func (z *Stop) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Stop{}, func(r rc_recipe.Recipe) {})
}

func (z *Stop) Preset() {
}
