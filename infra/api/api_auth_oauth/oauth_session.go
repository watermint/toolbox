package api_auth_oauth

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
	"sort"
)

func newOAuthAppWrapper(ctl app_control.Control, app api_auth.OAuthAppData) api_auth.OAuthApp {
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
	style := oauth2.AuthStyleAutoDetect
	switch z.app.EndpointStyle {
	case api_auth.AuthStyleInParams:
		style = oauth2.AuthStyleInParams
	case api_auth.AuthStyleInHeader:
		style = oauth2.AuthStyleInHeader
	}
	sortedScopes := make([]string, len(scope))
	copy(sortedScopes[:], scope[:])
	sort.Strings(sortedScopes)

	ak := api_appkey.New(z.ctl)
	clientId, clientSecret := ak.Key(z.app.AppKeyName)

	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   z.app.EndpointAuthUrl,
			TokenURL:  z.app.EndpointTokenUrl,
			AuthStyle: style,
		},
		RedirectURL: z.app.RedirectUrl,
		Scopes:      sortedScopes,
	}
}

func (z oauthAppWrapper) UsePKCE() bool {
	return z.app.UsePKCE
}

func NewSessionCodeAuth(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionCodeAuth{
		ctl: ctl,
		newConsole: func(ctl app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
			return NewConsoleOAuth(ctl, peerName, app)
		},
	}
}

func NewSessionRedirect(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionCodeAuth{
		ctl: ctl,
		newConsole: func(ctl app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
			return NewConsoleRedirect(ctl, peerName, app)
		},
	}
}

func NewSessionAlwaysFail(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionCodeAuth{
		ctl: ctl,
		newConsole: func(ctl app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
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
	newConsole func(ctl app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole
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

func NewSessionRepository(session api_auth.OAuthSession, repository api_auth.OAuthRepository) api_auth.OAuthSession {
	return &sessionRepository{
		session:    session,
		repository: repository,
	}
}

type sessionRepository struct {
	session    api_auth.OAuthSession
	repository api_auth.OAuthRepository
}

func (z sessionRepository) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	entity, err = z.session.Start(session)
	if err != nil {
		return api_auth.OAuthEntity{}, err
	}
	z.repository.Put(entity)
	return entity, err
}
