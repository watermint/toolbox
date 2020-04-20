package nw_diag

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	promptThreshold = 5 * 1000 * time.Millisecond
)

type MsgNetwork struct {
	ProgressTestingRemind app_msg.Message
	ProgressTesting       app_msg.Message
	ProgressTestingDone   app_msg.Message
	ErrorUnreachable      app_msg.Message
}

var (
	MNetwork        = app_msg.Apply(&MsgNetwork{}).(*MsgNetwork)
	NetworkDiagUrls = []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}
)

func Network(ctl app_control.Control) error {
	l := ctl.Log()
	ui := ctl.UI()
	ui.Info(MNetwork.ProgressTesting)
	doPrompt := true
	defer func() { doPrompt = false }()
	go func() {
		time.Sleep(promptThreshold)
		if doPrompt {
			ui.Info(MNetwork.ProgressTestingRemind)
		}
	}()

	for _, url := range NetworkDiagUrls {
		resp, err := http.Head(url)
		ll := l.With(zap.String("Url", url))
		if err != nil {
			ll.Debug("Network test failed", zap.Error(err))
			ui.Error(MNetwork.ErrorUnreachable.With("Url", url).With("Error", err))
			return err
		}

		if resp.StatusCode >= 400 {
			ll.Debug("Bad server response", zap.Int("status_code", resp.StatusCode))
			return errors.New("bad server response")
		}

		ll.Debug("Network test success", zap.Int("status_code", resp.StatusCode))
	}
	ui.Info(MNetwork.ProgressTestingDone)
	ui.Break()

	return nil
}
