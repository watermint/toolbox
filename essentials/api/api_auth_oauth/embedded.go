package api_auth_oauth

import (
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"reflect"
	"sort"
)

func NewSessionEmbedded(entity api_auth.OAuthEntity) api_auth.OAuthSession {
	return &sessionEmbedded{
		entity: entity,
	}
}

type sessionEmbedded struct {
	entity api_auth.OAuthEntity
}

func (z sessionEmbedded) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	if session.AppData.AppKeyName != z.entity.KeyName ||
		session.PeerName != z.entity.PeerName {
		return api_auth.OAuthEntity{}, errors.New("not found")
	}
	scopeSession := make([]string, len(session.Scopes))
	copy(scopeSession[:], session.Scopes[:])
	sort.Strings(scopeSession)
	scopeEntity := make([]string, len(z.entity.Scopes))
	copy(scopeEntity[:], z.entity.Scopes[:])
	sort.Strings(scopeEntity)
	if reflect.DeepEqual(scopeSession, scopeEntity) {
		return z.entity, nil
	}
	return api_auth.OAuthEntity{}, errors.New("not found")
}
