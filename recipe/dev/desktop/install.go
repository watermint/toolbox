package desktop

import (
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/log/es_process"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Install struct {
	rc_recipe.RemarkSecret
	InstallerUrl   string
	Silent         bool
	SilentNoLaunch bool
}

func (z *Install) Exec(c app_control.Control) error {
	l := c.Log()
	if runtime.GOOS != "windows" {
		l.Info("Skip: operation is not supported on this platform")
		return nil
	}
	dn := "DropboxOfflineInstaller.exe"
	dp := filepath.Join(c.Workspace().Job(), dn)

	arg := ""
	// https://help.dropbox.com/installs-integrations/desktop/enterprise-installer
	if z.Silent {
		arg = "/S"
	}
	if z.SilentNoLaunch {
		arg = "/NOLAUNCH"
	}

	if err := es_download.Download(c.Log(), z.InstallerUrl, dp); err != nil {
		l.Error("Unable to download installer", es_log.Error(err))
		return err
	}

	cmd := exec.Command(dp, arg)
	pl := es_process.NewLogger(cmd, c)
	pl.Start()
	defer pl.Close()

	l.Info("Start installer")
	if err := cmd.Start(); err != nil {
		l.Error("Unable to start installer", es_log.Error(err))
		return err
	}
	l.Info("Waiting for finish")
	if err := cmd.Wait(); err != nil {
		l.Error("Unable to wait", es_log.Error(err))
		return err
	}
	l.Info("Installation finished")
	return nil
}

func (z *Install) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}

func (z *Install) Preset() {
	z.InstallerUrl = "https://www.dropbox.com/download?full=1&os=win"
}
