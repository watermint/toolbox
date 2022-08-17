package api_auth_oauth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
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
	FailedOrCancelled   app_msg.Message
	ProgressAuthSuccess app_msg.Message
	OauthSeq1           app_msg.Message
	OauthSeq2           app_msg.Message
}

var (
	MApiAuth = app_msg.Apply(&MsgApiAuth{}).(*MsgApiAuth)
)

func NewConsoleOAuth(c app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
	return &OAuthConsole{
		ctl:      c,
		app:      app,
		peerName: peerName,
	}
}

type OAuthConsole struct {
	ctl      app_control.Control
	app      api_auth.OAuthApp
	peerName string
}

func (z *OAuthConsole) PeerName() string {
	return z.peerName
}

func (z *OAuthConsole) Start(scopes []string) (tc api_auth.OAuthContext, err error) {
	l := z.ctl.Log().With(esl.String("peerName", z.peerName), esl.Strings("scopes", scopes))
	ui := z.ctl.UI()

	l.Debug("Start OAuth sequence")
	t, err := z.oauthStart(scopes)
	if err != nil {
		l.Debug("Authentication finished with an error", esl.Error(err))
		ui.Error(MApiAuth.FailedOrCancelled.With("Cause", err))
		return nil, err
	}
	ui.Progress(MApiAuth.ProgressAuthSuccess)
	return api_auth.NewContext(t, z.app.Config(scopes), z.peerName, scopes), nil
}

func (z *OAuthConsole) oauthStart(scopes []string) (*oauth2.Token, error) {
	l := z.ctl.Log()
	l.Debug("Start OAuth sequence")
	state := sc_random.MustGetSecureRandomString(8)
	challenge := sc_random.MustGetSecureRandomString(64)

	tok, err := z.oauthAskCode(scopes, state, challenge)
	if err != nil {
		l.Debug("Authentication failed due to the error", esl.Error(err))
		return nil, err
	}
	return tok, nil
}

func (z *OAuthConsole) oauthUrl(cfg *oauth2.Config, state, challenge string) string {
	if z.app.UsePKCE() {
		challenge64 := base64.RawURLEncoding.EncodeToString([]byte(challenge))
		s256a32 := sha256.Sum256([]byte(challenge64))
		s256 := make([]byte, len(s256a32))
		copy(s256[:], s256a32[:])
		s256e := base64.RawURLEncoding.EncodeToString(s256)
		return cfg.AuthCodeURL(
			state,
			oauth2.SetAuthURLParam("code_challenge", s256e),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
			oauth2.SetAuthURLParam("response_type", "code"),
			oauth2.SetAuthURLParam("token_access_type", "offline"),
		)
	} else {
		return cfg.AuthCodeURL(
			state,
			oauth2.SetAuthURLParam("response_type", "code"),
		)
	}
}

func (z *OAuthConsole) oauthExchange(cfg *oauth2.Config, code, challenge string) (*oauth2.Token, error) {
	if z.app.UsePKCE() {
		challenge64 := base64.RawURLEncoding.EncodeToString([]byte(challenge))
		return cfg.Exchange(context.Background(),
			code,
			oauth2.SetAuthURLParam("code_verifier", challenge64),
		)
	} else {
		return cfg.Exchange(context.Background(), code)
	}
}

func (z *OAuthConsole) oauthCode(state, challenge string) string {
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

func (z *OAuthConsole) oauthAskCode(scopes []string, state, challenge string) (*oauth2.Token, error) {
	ui := z.ctl.UI()
	cfg := z.app.Config(scopes)
	url := z.oauthUrl(cfg, state, challenge)

	ui.Info(MApiAuth.OauthSeq1.With("Url", url))

	code := z.oauthCode(state, challenge)
	if code == "" {
		return nil, app.ErrorUserCancelled
	}

	return z.oauthExchange(cfg, code, challenge)
}
