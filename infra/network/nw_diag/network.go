package nw_diag

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
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
	l := ctl.WorkBundle().Summary().Logger()
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
		ll := l.With(esl.String("Url", url))
		if err != nil {
			ll.Debug("Network test failed", esl.Error(err))
			ui.Error(MNetwork.ErrorUnreachable.With("Url", url).With("Error", err))
			return err
		}

		if resp.StatusCode >= 400 {
			ll.Debug("Bad server response", esl.Int("status_code", resp.StatusCode))
			return errors.New("bad server response")
		}

		ll.Debug("Network test success", esl.Int("status_code", resp.StatusCode))
	}
	ui.Info(MNetwork.ProgressTestingDone)
	ui.Break()

	return nil
}
