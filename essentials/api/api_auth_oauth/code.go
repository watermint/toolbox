package api_auth_oauth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_apikey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"golang.org/x/oauth2"
	"strings"
	"time"
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

func NewSessionCodeAuth(ctl app_control.Control) api_auth2.OAuthSession {
	return &sessionCodeAuthImpl{
		ctl: ctl,
	}
}

// sessionCodeAuthImpl
type sessionCodeAuthImpl struct {
	ctl app_control.Control
}

func (z *sessionCodeAuthImpl) Start(session api_auth2.OAuthSessionData) (entity api_auth2.OAuthEntity, err error) {
	l := z.ctl.Log().With(esl.String("peerName", session.PeerName), esl.Strings("scopes", session.Scopes))
	ui := z.ctl.UI()

	l.Debug("Start OAuth sequence")
	token, err := z.oauthStart(session)
	if err != nil {
		l.Debug("Authentication finished with an error", esl.Error(err))
		ui.Error(MApiAuth.FailedOrCancelled.With("Cause", err))
		return api_auth2.NewNoAuthOAuthEntity(), err
	}
	ui.Progress(MApiAuth.ProgressAuthSuccess)
	return api_auth2.OAuthEntity{
		KeyName:  session.AppData.AppKeyName,
		Scopes:   session.Scopes,
		PeerName: session.PeerName,
		Token: api_auth2.OAuthTokenData{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			Expiry:       token.Expiry,
		},
		Description: "",
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}

func (z *sessionCodeAuthImpl) oauthStart(session api_auth2.OAuthSessionData) (*oauth2.Token, error) {
	l := z.ctl.Log()
	l.Debug("Start OAuth sequence")
	state := sc_random.MustGetSecureRandomString(8)
	challenge := sc_random.MustGetSecureRandomString(64)

	tok, err := z.oauthAskCode(session, state, challenge)
	if err != nil {
		l.Debug("Authentication failed due to the error", esl.Error(err))
		return nil, err
	}
	return tok, nil
}

func (z *sessionCodeAuthImpl) oauthUrl(session api_auth2.OAuthSessionData, cfg *oauth2.Config, state, challenge string) string {
	if session.AppData.UsePKCE {
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

func (z *sessionCodeAuthImpl) oauthExchange(session api_auth2.OAuthSessionData, cfg *oauth2.Config, code, challenge string) (*oauth2.Token, error) {
	if session.AppData.UsePKCE {
		challenge64 := base64.RawURLEncoding.EncodeToString([]byte(challenge))
		return cfg.Exchange(context.Background(),
			code,
			oauth2.SetAuthURLParam("code_verifier", challenge64),
		)
	} else {
		return cfg.Exchange(context.Background(), code)
	}
}

func (z *sessionCodeAuthImpl) oauthCode() string {
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

func (z *sessionCodeAuthImpl) oauthAskCode(session api_auth2.OAuthSessionData, state, challenge string) (*oauth2.Token, error) {
	ui := z.ctl.UI()
	cfg := session.AppData.Config(session.Scopes, func(appKey string) (clientId, clientSecret string) {
		return app_apikey.Resolve(z.ctl, session.AppData.AppKeyName)
	})
	url := z.oauthUrl(session, cfg, state, challenge)

	ui.Info(MApiAuth.OauthSeq1.With("Url", url))

	code := z.oauthCode()
	if code == "" {
		return nil, app.ErrorUserCancelled
	}

	return z.oauthExchange(session, cfg, code, challenge)
}
