package api_auth_oauth

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_repo"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
	"os"
	"reflect"
	"sort"
	"time"
)

var (
	ErrorNoExistingSession = errors.New("no existing session found")
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
		Timestamp:   time.Now().Format(time.RFC3339),
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
	return api_auth.NewNoAuthOAuthEntity(), ErrorNoExistingSession
}

func NewSessionDeployEnv(envName string) api_auth.OAuthSession {
	return &deployEnvSession{
		envName: envName,
	}
}

type deployEnvSession struct {
	envName string
}

func (z deployEnvSession) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	l := esl.Default()
	e := os.Getenv(z.envName)
	if e == "" {
		return api_auth.OAuthEntity{}, ErrorNoExistingSession
	}
	if err := json.Unmarshal([]byte(e), &entity); err != nil {
		l.Debug("Unable to unmarshal env", esl.Error(err))
		return api_auth.OAuthEntity{}, ErrorNoExistingSession
	}

	if entity.KeyName != session.AppData.AppKeyName {
		l.Debug("App Key does not mach", esl.String("expected", session.AppData.AppKeyName), esl.String("env", entity.KeyName))
		return api_auth.OAuthEntity{}, ErrorNoExistingSession
	}
	if entity.PeerName != session.PeerName {
		l.Debug("Peer name does not mach", esl.String("expected", session.PeerName), esl.String("env", entity.PeerName))
		return api_auth.OAuthEntity{}, ErrorNoExistingSession
	}
	entityScopes := make([]string, len(entity.Scopes))
	sessionScopes := make([]string, len(session.Scopes))
	copy(entityScopes[:], entity.Scopes[:])
	copy(sessionScopes[:], session.Scopes[:])
	sort.Strings(entityScopes)
	sort.Strings(sessionScopes)
	if !reflect.DeepEqual(entityScopes, sessionScopes) {
		l.Debug("Scope does not mach", esl.Strings("expected", sessionScopes), esl.Strings("env", entityScopes))
		return api_auth.OAuthEntity{}, ErrorNoExistingSession
	}

	return entity, nil
}
