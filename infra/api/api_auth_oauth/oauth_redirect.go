package api_auth_oauth

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_callback"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"golang.org/x/oauth2"
)

var (
	ErrorOAuthSequenceStopped = errors.New("the oauth sequence stopped")
	ErrorOAuthFailure         = errors.New("oauth failure")
	ErrorOAuthCancelled       = errors.New("auth cancelled")
)

type OptInFeatureRedirect struct {
	app_feature.OptInStatus
}

func NewConsoleRedirect(c app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
	return &Redirect{
		ctl:      c,
		app:      app,
		peerName: peerName,
	}
}

type Redirect struct {
	ctl      app_control.Control
	app      api_auth.OAuthApp
	peerName string
}

func (z *Redirect) PeerName() string {
	return z.peerName
}

func (z *Redirect) Start(scopes []string) (token api_auth.OAuthContext, err error) {
	l := z.ctl.Log().With(esl.Strings("scopes", scopes), esl.String("peerName", z.peerName))

	if z.ctl.Feature().IsTest() {
		return nil, qt_errors.ErrorSkipEndToEndTest
	}

	rs := &RedirectService{
		ctl:      z.ctl,
		app:      z.app,
		peerName: z.peerName,
		scopes:   scopes,
		state:    sc_random.MustGetSecureRandomString(8),
		result:   nil,
		token:    nil,
	}
	cb := api_callback.New(z.ctl, rs, app.DefaultWebPort)

	l.Debug("Starting sequence")
	if err := cb.Flow(); err != nil {
		l.Debug("Failure on the flow", esl.Error(err))
		return nil, err
	}

	done, result, err := rs.Result()
	if !done {
		l.Debug("RedirectService did not catch result")
		return nil, ErrorOAuthSequenceStopped
	}
	if !result {
		l.Debug("Auth failure", esl.Error(err))
		if err != nil {
			return nil, err
		} else {
			return nil, ErrorOAuthFailure
		}
	}

	t := rs.Token()
	if t == nil {
		l.Debug("No token available")
		return nil, ErrorOAuthFailure
	}

	l.Debug("Auth success")
	return api_auth.NewContext(t, z.app.Config(scopes), z.peerName, scopes), nil
}

type RedirectService struct {
	ctl         app_control.Control
	app         api_auth.OAuthApp
	peerName    string
	scopes      []string
	state       string
	result      *bool
	resultErr   error
	redirectUrl string
	token       *oauth2.Token
}

func (z *RedirectService) Token() *oauth2.Token {
	return z.token
}

func (z *RedirectService) Result() (done, result bool, err error) {
	if z.result == nil {
		return false, false, nil
	} else {
		return true, *z.result, z.resultErr
	}
}

func (z *RedirectService) Url(redirectUrl string) string {
	l := z.ctl.Log().With(esl.String("peerName", z.peerName), esl.Strings("scopes", z.scopes))
	cfg := z.app.Config(z.scopes)
	url := cfg.AuthCodeURL(
		z.state,
		oauth2.SetAuthURLParam("client_id", cfg.ClientID),
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("redirect_uri", redirectUrl),
	)
	l.Debug("generated url", esl.String("url", url))
	z.redirectUrl = redirectUrl
	return url
}

func (z *RedirectService) Verify(state, code string) bool {
	l := z.ctl.Log().With(esl.String("peerName", z.peerName), esl.Strings("scopes", z.scopes))

	if z.state != state {
		l.Debug("invalid state (csrf token)", esl.String("given", state), esl.String("expected", z.state))
		return false
	}

	cfg := z.app.Config(z.scopes)
	cfg.RedirectURL = z.redirectUrl
	l.Debug("exchange", esl.String("redirect", z.redirectUrl), esl.Any("cfg", cfg))
	token, err := cfg.Exchange(context.Background(), code)
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
