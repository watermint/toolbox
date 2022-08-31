package api_auth_oauth

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_repo"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
)

func newOAuthAppWrapper(ctl app_control.Control, app api_auth.OAuthAppData) api_auth.OAuthAppLegacy {
	return &oauthAppWrapper{
		app: app,
		ctl: ctl,
	}
}

type oauthAppWrapper struct {
	app api_auth.OAuthAppData
	ctl app_control.Control
}

func (z oauthAppWrapper) Config(scope []string) *oauth2.Config {
	return z.app.Config(scope, func(appKey string) (clientId, clientSecret string) {
		return api_appkey.Resolve(z.ctl, z.app.AppKeyName)
	})
}

func (z oauthAppWrapper) UsePKCE() bool {
	return z.app.UsePKCE
}

func NewSessionCodeAuth(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionCodeAuth{
		ctl: ctl,
		newConsole: func(ctl app_control.Control, peerName string, app api_auth.OAuthAppLegacy) api_auth.OAuthConsole {
			return NewConsoleOAuth(ctl, peerName, app)
		},
	}
}

func NewSessionRedirect(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionCodeAuth{
		ctl: ctl,
		newConsole: func(ctl app_control.Control, peerName string, app api_auth.OAuthAppLegacy) api_auth.OAuthConsole {
			return NewConsoleRedirect(ctl, peerName, app)
		},
	}
}

func NewSessionAlwaysFail(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionCodeAuth{
		ctl: ctl,
		newConsole: func(ctl app_control.Control, peerName string, app api_auth.OAuthAppLegacy) api_auth.OAuthConsole {
			return &sessionAlwaysFail{
				peerName: peerName,
			}
		},
	}
}

type sessionAlwaysFail struct {
	peerName string
}

func (z sessionAlwaysFail) PeerName() string {
	return z.peerName
}

func (z sessionAlwaysFail) Start(scope []string) (token api_auth.OAuthContext, err error) {
	return nil, errors.New("always fail")
}

type sessionCodeAuth struct {
	ctl        app_control.Control
	newConsole func(ctl app_control.Control, peerName string, app api_auth.OAuthAppLegacy) api_auth.OAuthConsole
}

func (z sessionCodeAuth) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	app := newOAuthAppWrapper(z.ctl, session.AppData)
	con := z.newConsole(z.ctl, session.PeerName, app)
	token, err := con.Start(session.Scopes)
	if err != nil {
		return api_auth.OAuthEntity{}, err
	}

	return api_auth.OAuthEntity{
		KeyName:  session.AppData.AppKeyName,
		Scopes:   session.Scopes,
		PeerName: session.PeerName,
		Token: api_auth.OAuthTokenData{
			AccessToken:  token.Token().AccessToken,
			RefreshToken: token.Token().RefreshToken,
			Expiry:       token.Token().Expiry,
		},
		Description: token.Description(),
	}, nil
}

func NewSessionRepository(session api_auth.OAuthSession, repository api_auth.Repository) api_auth.OAuthSession {
	return &sessionRepository{
		session:    session,
		repository: api_auth_repo.NewOAuth(repository),
	}
}

type sessionRepository struct {
	session    api_auth.OAuthSession
	repository api_auth.OAuthRepository
}

func (z sessionRepository) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	entity, found := z.repository.Get(session.AppData.AppKeyName, session.Scopes, session.PeerName)
	if found {
		return entity, nil
	}
	entity, err = z.session.Start(session)
	if err != nil {
		return api_auth.OAuthEntity{}, err
	}
	z.repository.Put(entity)
	return entity, nil
}

func NewSessionReadOnly(repository api_auth.Repository) api_auth.OAuthSession {
	return &readOnlySession{
		repository: api_auth_repo.NewOAuth(repository),
	}
}

type readOnlySession struct {
	repository api_auth.OAuthRepository
}

func (z readOnlySession) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	entity, found := z.repository.Get(session.AppData.AppKeyName, session.Scopes, session.PeerName)
	if found {
		return entity, nil
	}
	return api_auth.NewNoAuthOAuthEntity(), errors.New("no existing token")
}
