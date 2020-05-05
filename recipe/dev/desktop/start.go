package desktop

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/log/es_process"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"os/exec"
	"runtime"
)

type Start struct {
	rc_recipe.RemarkSecret
}

func (z *Start) Exec(c app_control.Control) error {
	l := c.Log()
	desktopAppPath := ""
	switch runtime.GOOS {
	case "windows":
		desktopAppPath = "C:/Program Files (x86)/Dropbox/Client/Dropbox.exe"

	case "darwin":
		desktopAppPath = "/Applications/Dropbox.app/Contents/MacOS/Dropbox"

	default:
		c.Log().Info("Skip: the command is not supported on this platform")
		return nil
	}

	cmd := exec.Command(desktopAppPath, "/home")
	pl := es_process.NewLogger(cmd, c)
	pl.Start()
	defer pl.Close()

	l.Info("Start Dropbox")
	err := cmd.Start()
	if err != nil {
		l.Error("Unable to start Desktop", es_log.Error(err))
		return err
	}

	l.Info("Waiting for Dropbox startup")
	cmd.Wait()
	return nil
}

func (z *Start) Test(c app_control.Control) error {
	if qt_endtoend.IsSkipEndToEndTest() {
		return nil
	}
	return rc_exec.Exec(c, &Start{}, func(r rc_recipe.Recipe) {})
}

func (z *Start) Preset() {
}
