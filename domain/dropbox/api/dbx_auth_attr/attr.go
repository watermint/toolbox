package dbx_auth_attr

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgAttr struct {
	ErrorVerifyFailed app_msg.Message
	AttrTeamLicenses  app_msg.Message
}

var (
	MAttr = app_msg.Apply(&MsgAttr{}).(*MsgAttr)
)

type Attr struct {
	app  api_auth.App
	ctl  app_control.Control
	auth api_auth.Console
}

func (z *Attr) PeerName() string {
	return z.auth.PeerName()
}

func (z *Attr) Auth(scopes []string) (token api_auth.Context, err error) {
	l := z.ctl.Log().With(esl.String("peerName", z.PeerName()), esl.Strings("scopes", scopes))
	ui := z.ctl.UI()

	tc, err := z.auth.Auth(scopes)
	if err != nil {
		return nil, err
	}

	l.Debug("Start verify token")

	tc, err = VerifyToken(z.PeerName(), tc, z.ctl, z.app)
	if err != nil {
		l.Debug("failed verify token", esl.Error(err))
		ui.Error(MAttr.ErrorVerifyFailed.With("Error", err))
		return nil, err
	}
	return tc, nil
}
