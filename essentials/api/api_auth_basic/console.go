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

type consoleOpts struct {
	CustomAskUserName app_msg.Message
	CustomAskPassword app_msg.Message
}

type ConsoleOpt func(o consoleOpts) consoleOpts

func CustomAskUserName(message app_msg.Message) ConsoleOpt {
	return func(o consoleOpts) consoleOpts {
		o.CustomAskUserName = message
		return o
	}
}
func CustomAskPassword(message app_msg.Message) ConsoleOpt {
	return func(o consoleOpts) consoleOpts {
		o.CustomAskPassword = message
		return o
	}
}

func (z consoleOpts) Apply(opts []ConsoleOpt) consoleOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

func NewConsole(ctl app_control.Control, opts ...ConsoleOpt) api_auth.BasicSession {
	return &consoleImpl{
		ctl:  ctl,
		opts: consoleOpts{}.Apply(opts),
	}
}

type consoleImpl struct {
	ctl  app_control.Control
	opts consoleOpts
}

func (z consoleImpl) Start(session api_auth.BasicSessionData) (entity api_auth.BasicEntity, err error) {
	ui := z.ctl.UI()
	ui.Info(MConsole.PromptEnterUsernameAndPassword)
	credential := api_auth.BasicCredential{}
	var cancel bool
	if !session.AppData.DontUseUsername {
		askMsg := MConsole.AskUserName
		if z.opts.CustomAskUserName != nil {
			askMsg = z.opts.CustomAskUserName
		}
		credential.Username, cancel = ui.AskText(askMsg)
		if cancel {
			return api_auth.BasicEntity{}, ErrorUserCancelled
		}
	}
	if !session.AppData.DontUsePassword {
		askMsg := MConsole.AskPassword
		if z.opts.CustomAskPassword != nil {
			askMsg = z.opts.CustomAskPassword
		}
		credential.Password, cancel = ui.AskSecure(askMsg)
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
