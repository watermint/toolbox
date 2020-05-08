package api_auth_impl

import (
	"context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"golang.org/x/oauth2"
	"strings"
)

type MsgApiAuth struct {
	FailedOrCancelled app_msg.Message
	OauthSeq1         app_msg.Message
	OauthSeq2         app_msg.Message
}

var (
	MApiAuth = app_msg.Apply(&MsgApiAuth{}).(*MsgApiAuth)
)

func NewConsoleOAuth(c app_control.Control, peerName string) api_auth.Console {
	return &Console{
		ctl:      c,
		app:      dbx_auth.NewApp(c),
		peerName: peerName,
	}
}

type Console struct {
	ctl      app_control.Control
	app      api_auth.App
	peerName string
}

func (z *Console) PeerName() string {
	return z.peerName
}

func (z *Console) Auth(scope string) (tc api_auth.Context, err error) {
	l := z.ctl.Log().With(esl.String("peerName", z.peerName), esl.String("scope", scope))
	ui := z.ctl.UI()

	l.Debug("Start OAuth sequence")
	t, err := z.oauthStart(scope)
	if err != nil {
		l.Debug("Authentication finished with an error", esl.Error(err))
		ui.Error(MApiAuth.FailedOrCancelled.With("Cause", err))
		return nil, err
	}
	return api_auth.NewContext(t, z.peerName, scope), nil
}

func (z *Console) oauthStart(scope string) (*oauth2.Token, error) {
	l := z.ctl.Log()
	l.Debug("Start OAuth sequence")
	state, err := sc_random.GenerateRandomString(8)
	if err != nil {
		l.Error("Unable to generate `state`", esl.Error(err))
		return nil, err
	}

	tok, err := z.oauthAskCode(scope, state)
	if err != nil {
		l.Debug("Authentication failed due to the error", esl.Error(err))
		return nil, err
	}
	return tok, nil
}

func (z *Console) oauthUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (z *Console) oauthExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (z *Console) oauthCode(state string) string {
	ui := z.ctl.UI()
	for {
		code, cancel := ui.AskText(MApiAuth.OauthSeq2)
		if cancel {
			return ""
		}
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim
		}
	}
}

func (z *Console) oauthAskCode(tokenType, state string) (*oauth2.Token, error) {
	ui := z.ctl.UI()
	cfg := z.app.Config(tokenType)
	url := z.oauthUrl(cfg, state)

	ui.Info(MApiAuth.OauthSeq1.With("Url", url))

	code := z.oauthCode(state)
	if code == "" {
		return nil, app.ErrorUserCancelled
	}

	return z.oauthExchange(cfg, code)
}
