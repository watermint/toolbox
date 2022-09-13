package api_auth_basic

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"time"
)

var (
	ErrorUserCancelled = app.ErrorUserCancelled
)

type MsgConsole struct {
	PromptEnterUsernameAndPassword app_msg.Message
	AskUserName                    app_msg.Message
	AskPassword                    app_msg.Message
}

var (
	MConsole = app_msg.Apply(&MsgConsole{}).(*MsgConsole)
)

func NewConsole(ctl app_control.Control) api_auth.BasicSession {
	return &consoleImpl{
		ctl: ctl,
	}
}

type consoleImpl struct {
	ctl app_control.Control
}

func (z consoleImpl) Start(session api_auth.BasicSessionData) (entity api_auth.BasicEntity, err error) {
	ui := z.ctl.UI()
	ui.Info(MConsole.PromptEnterUsernameAndPassword)
	credential := api_auth.BasicCredential{}
	var cancel bool
	if !session.AppData.DontUseUsername {
		credential.Username, cancel = ui.AskSecure(MConsole.AskUserName)
		if cancel {
			return api_auth.BasicEntity{}, ErrorUserCancelled
		}
	}
	if !session.AppData.DontUsePassword {
		credential.Password, cancel = ui.AskSecure(MConsole.AskPassword)
		if cancel {
			return api_auth.BasicEntity{}, ErrorUserCancelled
		}
	}
	return api_auth.BasicEntity{
		KeyName:     session.AppData.AppKeyName,
		PeerName:    session.PeerName,
		Credential:  credential,
		Description: "",
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}
