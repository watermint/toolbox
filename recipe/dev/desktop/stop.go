package desktop

import (
	"github.com/andybrewer/mack"
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_process"
	"go.uber.org/zap"
	"math"
	"os/exec"
	"runtime"
	"time"
)

// Tell Dropbox to quit, but no guarantee of stop the process.
type Stop struct {
	rc_recipe.RemarkSecret
	WaitSeconds mo_int.RangeInt
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
	if z.WaitSeconds.Value() > 0 {
		c.Log().Info("Waiting for stop", zap.Int("seconds", int(z.WaitSeconds.Value())))
		time.Sleep(time.Duration(z.WaitSeconds.Value()) * time.Second)
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
	return rc_exec.Exec(c, &Stop{}, func(r rc_recipe.Recipe) {
		m := r.(*Stop)
		m.WaitSeconds.SetValue(0)
	})
}

func (z *Stop) Preset() {
	z.WaitSeconds.SetRange(0, math.MaxInt32, 0)
}
