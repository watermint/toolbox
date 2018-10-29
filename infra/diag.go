package infra

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/user"
	"runtime"
)

type Diag struct {
	ExecContext *ExecContext
}

func (d *Diag) Log() *zap.Logger {
	return d.ExecContext.Log()
}

func (d *Diag) Runtime() error {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	usr, _ := user.Current()

	d.Log().Debug(
		"Runtime",
		zap.String("os", runtime.GOOS),
		zap.String("arch", runtime.GOARCH),
		zap.String("go_version", runtime.Version()),
		zap.Int("num_cpu", runtime.NumCPU()),
		zap.String("hostname", hostname),
		zap.String("pwd", wd),
		zap.String("user_uid", usr.Uid),
		zap.String("user_name", usr.Name),
		zap.String("user_home", usr.HomeDir),
		zap.Int("pid", os.Getpid()),
		zap.Int("euid", os.Geteuid()),
		zap.Strings("env", os.Environ()),
	)
	return nil
}

func (d *Diag) Network() error {
	urls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}

	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			d.Log().Debug(
				"Network test failed",
				zap.String("url", url),
				zap.Error(err),
			)
			return err
		}

		if resp.StatusCode >= 400 {
			d.Log().Debug(
				"Bad server response",
				zap.String("url", url),
				zap.Int("status_code", resp.StatusCode),
			)
			return errors.New("bad server response")
		}

		d.Log().Debug(
			"Network test success",
			zap.String("url", url),
			zap.Int("status_code", resp.StatusCode),
		)
	}
	return nil
}
