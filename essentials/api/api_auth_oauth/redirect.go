package api_auth_oauth

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_callback"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_apikey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"golang.org/x/oauth2"
	"time"
)

var (
	ErrorOAuthSequenceStopped = errors.New("the oauth sequence stopped")
	ErrorOAuthFailure         = errors.New("oauth failure")
	ErrorOAuthCancelled       = errors.New("auth cancelled")
)

type OptInFeatureRedirect struct {
	app_feature.OptInStatus
}

func NewSessionRedirect(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionRedirectImpl{
		ctl: ctl,
	}
}

type sessionRedirectImpl struct {
	ctl app_control.Control
}

func (z *sessionRedirectImpl) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	l := z.ctl.Log().With(esl.Strings("scopes", session.Scopes), esl.String("peerName", session.PeerName))

	if z.ctl.Feature().IsTest() {
		return api_auth.NewNoAuthOAuthEntity(), qt_errors.ErrorSkipEndToEndTest
	}

	cfg := session.AppData.Config(session.Scopes, func(appKey string) (clientId, clientSecret string) {
		return app_apikey.Resolve(z.ctl, session.AppData.AppKeyName)
	})
	rs := &redirectServiceImpl{
		ctl:     z.ctl,
		cfg:     cfg,
		session: session,
		state:   sc_random.MustGetSecureRandomString(8),
		result:  nil,
		token:   nil,
	}
	cb := api_callback.New(z.ctl, rs, app_definitions.DefaultWebPort, session.UseSecureRedirect)

	l.Debug("Starting sequence")
	if err := cb.Flow(); err != nil {
		l.Debug("Failure on the flow", esl.Error(err))
		return api_auth.NewNoAuthOAuthEntity(), err
	}

	done, result, err := rs.Result()
	if !done {
		l.Debug("redirectServiceImpl did not catch result")
		return api_auth.NewNoAuthOAuthEntity(), ErrorOAuthSequenceStopped
	}
	if !result {
		l.Debug("Auth failure", esl.Error(err))
		if err != nil {
			return api_auth.NewNoAuthOAuthEntity(), err
		} else {
			return api_auth.NewNoAuthOAuthEntity(), ErrorOAuthFailure
		}
	}

	t := rs.Token()
	if t == nil {
		l.Debug("No token available")
		return api_auth.NewNoAuthOAuthEntity(), ErrorOAuthFailure
	}

	l.Debug("Auth success")
	return api_auth.OAuthEntity{
		KeyName:  session.AppData.AppKeyName,
		Scopes:   session.Scopes,
		PeerName: session.PeerName,
		Token: api_auth.OAuthTokenData{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			Expiry:       t.Expiry,
		},
		Description: "",
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}

type redirectServiceImpl struct {
	ctl         app_control.Control
	cfg         *oauth2.Config
	session     api_auth.OAuthSessionData
	state       string
	result      *bool
	resultErr   error
	redirectUrl string
	token       *oauth2.Token
}

func (z *redirectServiceImpl) Token() *oauth2.Token {
	return z.token
}

func (z *redirectServiceImpl) Result() (done, result bool, err error) {
	if z.result == nil {
		return false, false, nil
	} else {
		return true, *z.result, z.resultErr
	}
}

func (z *redirectServiceImpl) Url(redirectUrl string) string {
	l := z.ctl.Log().With(esl.String("peerName", z.session.PeerName), esl.Strings("scopes", z.session.Scopes))

	url := z.cfg.AuthCodeURL(
		z.state,
		oauth2.SetAuthURLParam("client_id", z.cfg.ClientID),
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("redirect_uri", redirectUrl),
	)
	l.Debug("generated url", esl.String("url", url))
	z.redirectUrl = redirectUrl
	return url
}

func (z *redirectServiceImpl) Verify(state, code string) bool {
	l := z.ctl.Log().With(esl.String("peerName", z.session.PeerName), esl.Strings("scopes", z.session.Scopes))

	if z.state != state {
		l.Debug("invalid state (csrf token)", esl.String("given", state), esl.String("expected", z.state))
		return false
	}

	z.cfg.RedirectURL = z.redirectUrl
	l.Debug("exchange", esl.String("redirect", z.redirectUrl), esl.Any("cfg", z.cfg))
	token, err := z.cfg.Exchange(context.Background(), code)
	if err != nil {
		l.Debug("Verification failure", esl.Error(err))
		t := false
		z.token = nil
		z.result = &t
		z.resultErr = err
		return false
	}

	l.Debug("Verification success")
	t := true
	z.token = token
	z.result = &t
	return true
}
