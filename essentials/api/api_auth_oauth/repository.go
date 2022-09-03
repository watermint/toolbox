package api_auth_oauth

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
)

func NewSessionRepository(session api_auth2.OAuthSession, repository api_auth2.Repository) api_auth2.OAuthSession {
	return &sessionRepository{
		session:    session,
		repository: api_auth_repo.NewOAuth(repository),
	}
}

type sessionRepository struct {
	session    api_auth2.OAuthSession
	repository api_auth2.OAuthRepository
}

func (z sessionRepository) Start(session api_auth2.OAuthSessionData) (entity api_auth2.OAuthEntity, err error) {
	entity, found := z.repository.Get(session.AppData.AppKeyName, session.Scopes, session.PeerName)
	if found {
		return entity, nil
	}
	entity, err = z.session.Start(session)
	if err != nil {
		return api_auth2.OAuthEntity{}, err
	}
	z.repository.Put(entity)
	return entity, nil
}
