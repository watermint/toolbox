package api_auth_oauth

import (
	"encoding/json"
	"errors"
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"reflect"
	"sort"
)

var (
	ErrorNoExistingSession = errors.New("no existing session found")
)

func NewSessionDeployEnv(envName string) api_auth2.OAuthSession {
	return &deployEnvSession{
		envName: envName,
	}
}

type deployEnvSession struct {
	envName string
}

func (z deployEnvSession) Start(session api_auth2.OAuthSessionData) (entity api_auth2.OAuthEntity, err error) {
	l := esl.Default()
	e := os.Getenv(z.envName)
	if e == "" {
		return api_auth2.OAuthEntity{}, ErrorNoExistingSession
	}
	if err := json.Unmarshal([]byte(e), &entity); err != nil {
		l.Debug("Unable to unmarshal env", esl.Error(err))
		return api_auth2.OAuthEntity{}, ErrorNoExistingSession
	}

	if entity.KeyName != session.AppData.AppKeyName {
		l.Debug("App Key does not mach", esl.String("expected", session.AppData.AppKeyName), esl.String("env", entity.KeyName))
		return api_auth2.OAuthEntity{}, ErrorNoExistingSession
	}
	if entity.PeerName != session.PeerName {
		l.Debug("Peer name does not mach", esl.String("expected", session.PeerName), esl.String("env", entity.PeerName))
		return api_auth2.OAuthEntity{}, ErrorNoExistingSession
	}
	entityScopes := make([]string, len(entity.Scopes))
	sessionScopes := make([]string, len(session.Scopes))
	copy(entityScopes[:], entity.Scopes[:])
	copy(sessionScopes[:], session.Scopes[:])
	sort.Strings(entityScopes)
	sort.Strings(sessionScopes)
	if !reflect.DeepEqual(entityScopes, sessionScopes) {
		l.Debug("Scope does not mach", esl.Strings("expected", sessionScopes), esl.Strings("env", entityScopes))
		return api_auth2.OAuthEntity{}, ErrorNoExistingSession
	}

	return entity, nil
}
