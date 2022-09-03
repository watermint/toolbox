package api_auth_oauth

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
)

func NewSessionReadOnly(repository api_auth2.Repository) api_auth2.OAuthSession {
	return &readOnlySession{
		repository: api_auth_repo.NewOAuth(repository),
	}
}

type readOnlySession struct {
	repository api_auth2.OAuthRepository
}

func (z readOnlySession) Start(session api_auth2.OAuthSessionData) (entity api_auth2.OAuthEntity, err error) {
	entity, found := z.repository.Get(session.AppData.AppKeyName, session.Scopes, session.PeerName)
	if found {
		return entity, nil
	}
	return api_auth2.NewNoAuthOAuthEntity(), ErrorNoExistingSession
}
