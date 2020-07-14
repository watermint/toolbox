package dbx_auth

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"golang.org/x/oauth2"
	"strings"
)

type MsgGenerated struct {
	GeneratedToken1 app_msg.Message
	GeneratedToken2 app_msg.Message
}

var (
	MGenerated = app_msg.Apply(&MsgGenerated{}).(*MsgGenerated)
)

func NewConsoleGenerated(c app_control.Control, peerName string) api_auth.Console {
	return &Generated{
		ctl:      c,
		peerName: peerName,
	}
}

type Generated struct {
	ctl      app_control.Control
	peerName string
}

func (z *Generated) PeerName() string {
	return z.peerName
}

func (z *Generated) Auth(scopes []string) (tc api_auth.Context, err error) {
	token, err := z.generatedToken(scopes)
	if err != nil {
		return nil, err
	}
	return api_auth.NewContext(token, z.peerName, scopes), nil
}

func (z *Generated) generatedTokenInstruction(scope string) {
	ui := z.ctl.UI()
	api := ""
	toa := ""

	switch scope {
	case api_auth.DropboxTokenFull:
		api = "Dropbox API"
		toa = "Full Dropbox"
	case api_auth.DropboxTokenApp:
		api = "Dropbox API"
		toa = "App folder"
	case api_auth.DropboxTokenBusinessInfo:
		api = "Dropbox Business API"
		toa = "Team information"
	case api_auth.DropboxTokenBusinessAudit:
		api = "Dropbox Business API"
		toa = "Team auditing"
	case api_auth.DropboxTokenBusinessFile:
		api = "Dropbox Business API"
		toa = "Team member file access"
	case api_auth.DropboxTokenBusinessManagement:
		api = "Dropbox Business API"
		toa = "Team member management"
	default:
		z.ctl.Log().Error("Undefined token type", esl.String("type", scope))
	}

	ui.Info(MGenerated.GeneratedToken1.With("API", api).With("TypeOfAccess", toa))
}

func (z *Generated) generatedToken(scopes []string) (*oauth2.Token, error) {
	if len(scopes) != 1 {
		esl.Default().Warn("Unsupported scopes", esl.Strings("scopes", scopes))
		return nil, errors.New("unsupported scope")
	}
	scope := scopes[0]
	ui := z.ctl.UI()
	z.generatedTokenInstruction(scope)
	for {
		code, cancel := ui.AskSecure(MGenerated.GeneratedToken2)
		if cancel {
			return nil, app.ErrorUserCancelled
		}
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return &oauth2.Token{AccessToken: trim}, nil
		}
	}
}
