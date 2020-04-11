package gh_auth

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_callback"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"strings"
)

const (
	// Read-only access to public information.
	ScopeNoScope = ""

	// full access to private and public repositories.
	ScopeRepo = "repo"

	// Read write access to public and private repository commit statuses.
	ScopeRepoStatus = "repo:status"

	// Deployment statuses for public and private repositories.
	ScopeRepoDeployment = "repo_deployment"

	// Read/write access to code, commit statuses, repository etc for public repositories.
	ScopePublicRepo = "public_repo"
)

type App struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z *App) AppKey(scope string) (key, secret string) {
	return z.res.Key(scope)
}

func (z *App) Config(scope string) *oauth2.Config {
	key, secret := z.AppKey(api_auth.Github)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint:     github.Endpoint,
		Scopes:       strings.Split(scope, ","),
	}
}

// TODO: Copy from Dropbox module. It could share most code.

func New(c app_control.Control, peerName string) api_auth.Console {
	return &Github{
		ctl:      c,
		app:      &App{c, api_appkey.New(c)},
		peerName: peerName,
	}
}

type Github struct {
	ctl      app_control.Control
	app      api_auth.App
	peerName string
}

func (z *Github) PeerName() string {
	return z.peerName
}

var (
	ErrorOAuthSequenceStopped = errors.New("the oauth sequence stopped")
	ErrorOAuthFailure         = errors.New("")
)

func (z *Github) Auth(scope string) (token api_auth.Context, err error) {
	l := z.ctl.Log().With(zap.String("scope", scope), zap.String("peerName", z.peerName))
	rs := &GithubService{
		ctl:      z.ctl,
		app:      z.app,
		peerName: z.peerName,
		scope:    scope,
		state:    sc_random.MustGenerateRandomString(8),
		result:   nil,
		token:    nil,
	}
	cb := api_callback.New(z.ctl, rs, app.DefaultWebPort)

	l.Debug("Starting sequence")
	if err := cb.Flow(); err != nil {
		l.Debug("Failure on the flow", zap.Error(err))
		return nil, err
	}

	done, result, err := rs.Result()
	if !done {
		l.Debug("RedirectService did not catch result")
		return nil, ErrorOAuthSequenceStopped
	}
	if !result {
		l.Debug("Auth failure", zap.Error(err))
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
	return api_auth.NewContext(t, z.peerName, scope), nil
}

type GithubService struct {
	ctl         app_control.Control
	app         api_auth.App
	peerName    string
	scope       string
	state       string
	result      *bool
	resultErr   error
	redirectUrl string
	token       *oauth2.Token
}

func (z *GithubService) Token() *oauth2.Token {
	return z.token
}

func (z *GithubService) Result() (done, result bool, err error) {
	if z.result == nil {
		return false, false, nil
	} else {
		return true, *z.result, z.resultErr
	}
}

func (z *GithubService) Url(redirectUrl string) string {
	l := z.ctl.Log().With(zap.String("peerName", z.peerName), zap.String("scope", z.scope))
	cfg := z.app.Config(z.scope)
	url := cfg.AuthCodeURL(
		z.state,
		oauth2.SetAuthURLParam("client_id", cfg.ClientID),
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("redirect_uri", redirectUrl),
	)
	l.Debug("generated url", zap.String("url", url))
	z.redirectUrl = redirectUrl
	return url
}

func (z *GithubService) Verify(state, code string) bool {
	l := z.ctl.Log().With(zap.String("peerName", z.peerName), zap.String("scope", z.scope))

	if z.state != state {
		l.Debug("invalid state (csrf token)", zap.String("given", state), zap.String("expected", z.state))
		return false
	}

	cfg := z.app.Config(z.scope)
	cfg.RedirectURL = z.redirectUrl
	l.Debug("exchange", zap.String("redirect", z.redirectUrl), zap.Any("cfg", cfg))
	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		l.Debug("Verification failure", zap.Error(err))
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
