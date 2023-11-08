package api_auth_key

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"time"
)

type MsgConsole struct {
	PromptEnterKey app_msg.Message
	AskKey         app_msg.Message
}

var (
	MConsole = app_msg.Apply(&MsgConsole{}).(*MsgConsole)
)

// NewConsole Create a new console session
// If askKeyPrompt is nil, use default prompt message.
func NewConsole(ctl app_control.Control, askKeyPrompt app_msg.Message) api_auth.KeySession {
	return &consoleImpl{
		ctl:    ctl,
		prompt: askKeyPrompt,
	}
}

type consoleImpl struct {
	ctl    app_control.Control
	prompt app_msg.Message
}

func (z consoleImpl) Start(session api_auth.KeySessionData) (entity api_auth.KeyEntity, err error) {
	ui := z.ctl.UI()
	ui.Info(MConsole.PromptEnterKey)
	var ask app_msg.Message
	var cancel bool
	if z.prompt != nil {
		ask = z.prompt
	} else {
		ask = MConsole.AskKey
	}
	credential := api_auth.KeyCredential{}
	credential.Key, cancel = ui.AskSecure(ask)

	if cancel {
		return api_auth.KeyEntity{}, app.ErrorUserCancelled
	}
	return api_auth.KeyEntity{
		KeyName:     session.AppData.AppKeyName,
		PeerName:    session.PeerName,
		Credential:  credential,
		Description: "",
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}
