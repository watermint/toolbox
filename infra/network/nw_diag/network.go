package nw_diag

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"net/http"
)

func Network(ctl app_control.Control) error {
	urls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}
	l := ctl.Log()
	ui := ctl.UI()
	ui.InfoK("run.network.progress.testing")

	for _, url := range urls {
		resp, err := http.Head(url)
		ll := l.With(zap.String("Url", url))
		if err != nil {
			ll.Debug("Network test failed", zap.Error(err))
			ui.ErrorK("run.network.error.unreachable",
				app_msg.P{
					"Url":   url,
					"Error": err,
				},
			)
			return err
		}

		if resp.StatusCode >= 400 {
			ll.Debug("Bad server response", zap.Int("status_code", resp.StatusCode))
			return errors.New("bad server response")
		}

		ll.Debug("Network test success", zap.Int("status_code", resp.StatusCode))
	}
	ui.InfoK("run.network.progress.testing.done")
	ui.Break()

	return nil
}
