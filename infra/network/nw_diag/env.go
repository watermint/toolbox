package nw_diag

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"os"
	"os/user"
	"runtime"
)

func Runtime(ctl app_control.Control) error {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	usr, _ := user.Current()
	l := ctl.WorkBundle().Summary().Logger()

	l.Debug("Runtime",
		es_log.String("os", runtime.GOOS),
		es_log.String("arch", runtime.GOARCH),
		es_log.String("go_version", runtime.Version()),
		es_log.Int("num_cpu", runtime.NumCPU()),
	)
	l.Debug("User",
		es_log.String("user_uid", usr.Uid),
		es_log.String("user_name", usr.Name),
		es_log.String("user_home", usr.HomeDir),
		es_log.String("pwd", wd),
	)
	l.Debug("Environment", es_log.Strings("env", os.Environ()))
	l.Debug("Process", es_log.Int("pid", os.Getpid()), es_log.Int("euid", os.Geteuid()))
	l.Debug("Hostname", es_log.String("hostname", hostname))
	l.Debug("Arguments", es_log.Strings("arg", os.Args))

	return nil
}
