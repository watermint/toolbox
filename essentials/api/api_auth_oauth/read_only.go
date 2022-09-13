package api_auth_oauth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
)

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
