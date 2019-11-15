package nw_diag

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"os"
	"os/user"
	"runtime"
)

func Runtime(ctl app_control.Control) error {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	usr, _ := user.Current()
	l := ctl.Log()

	l.Debug("Runtime",
		zap.String("os", runtime.GOOS),
		zap.String("arch", runtime.GOARCH),
		zap.String("go_version", runtime.Version()),
		zap.Int("num_cpu", runtime.NumCPU()),
	)
	l.Debug("User",
		zap.String("user_uid", usr.Uid),
		zap.String("user_name", usr.Name),
		zap.String("user_home", usr.HomeDir),
		zap.String("pwd", wd),
	)
	l.Debug("Environment", zap.Strings("env", os.Environ()))
	l.Debug("Process", zap.Int("pid", os.Getpid()), zap.Int("euid", os.Geteuid()))
	l.Debug("Hostname", zap.String("hostname", hostname))
	l.Debug("Arguments", zap.Strings("arg", os.Args))

	return nil
}
