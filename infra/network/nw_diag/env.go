package nw_diag

import (
	"github.com/watermint/toolbox/essentials/log/esl"
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
		esl.String("os", runtime.GOOS),
		esl.String("arch", runtime.GOARCH),
		esl.String("go_version", runtime.Version()),
		esl.Int("num_cpu", runtime.NumCPU()),
	)
	l.Debug("User",
		esl.String("user_uid", usr.Uid),
		esl.String("user_name", usr.Name),
		esl.String("user_home", usr.HomeDir),
		esl.String("pwd", wd),
	)
	l.Debug("Environment", esl.Strings("env", os.Environ()))
	l.Debug("Process", esl.Int("pid", os.Getpid()), esl.Int("euid", os.Geteuid()))
	l.Debug("Hostname", esl.String("hostname", hostname))
	l.Debug("Arguments", esl.Strings("arg", os.Args))

	return nil
}
