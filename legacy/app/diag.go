package app

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

func (z *Diag) Log() *zap.Logger {
	return z.ExecContext.Log()
}

func (z *Diag) Runtime() error {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	usr, _ := user.Current()

	z.Log().Debug(
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
	z.Log().Debug("Command", zap.Strings("arg", os.Args))
	return nil
}

func (z *Diag) Network() error {
	urls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}

	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			z.Log().Debug(
				"Network test failed",
				zap.String("url", url),
				zap.Error(err),
			)
			z.ExecContext.Msg("app.common.diag.network.err.unreachable").WithData(struct {
				Url   string
				Error string
			}{
				Url:   url,
				Error: err.Error(),
			}).TellError()

			return err
		}

		if resp.StatusCode >= 400 {
			z.Log().Debug(
				"Bad server response",
				zap.String("url", url),
				zap.Int("status_code", resp.StatusCode),
			)
			return errors.New("bad server response")
		}

		z.Log().Debug(
			"Network test success",
			zap.String("url", url),
			zap.Int("status_code", resp.StatusCode),
		)
	}
	return nil
}
