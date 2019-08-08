package app_diag

import (
	"errors"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_msg"
	"go.uber.org/zap"
	"net/http"
)

func Network(ctl app_control.Control) error {
	urls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}
	l := ctl.Log()

	ctl.UI().Info("run.network.progress.testing")

	for _, url := range urls {
		resp, err := http.Head(url)
		ll := l.With(zap.String("Url", url))
		if err != nil {
			ll.Debug("Network test failed", zap.Error(err))
			ctl.UI().Error("run.network.error.unreachable",
				app_msg.P("Url", url),
				app_msg.P("Error", err),
			)
			return err
		}

		if resp.StatusCode >= 400 {
			ll.Debug("Bad server response", zap.Int("status_code", resp.StatusCode))
			return errors.New("bad server response")
		}

		ll.Debug("Network test success", zap.Int("status_code", resp.StatusCode))
	}
	ctl.UI().Info("run.network.progress.testing.done")
	ctl.UI().Break()

	return nil
}
