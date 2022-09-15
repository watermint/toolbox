package api_auth_basic

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
)

func NewRepository(session api_auth.BasicSession, repository api_auth.Repository) api_auth.BasicSession {
	return &repoImpl{
		session:    session,
		repository: api_auth_repo.NewBasic(repository),
	}
}

type repoImpl struct {
	session    api_auth.BasicSession
	repository api_auth.BasicRepository
}

func (z repoImpl) Start(session api_auth.BasicSessionData) (entity api_auth.BasicEntity, err error) {
	entity, found := z.repository.Get(session.AppData.AppKeyName, session.PeerName)
	if found {
		return entity, nil
	}
	entity, err = z.session.Start(session)
	if err != nil {
		return api_auth.BasicEntity{}, err
	}
	z.repository.Put(entity)
	return entity, nil
}
