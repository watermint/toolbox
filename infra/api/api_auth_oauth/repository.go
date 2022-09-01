package api_auth_oauth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_repo"
)

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
