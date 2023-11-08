package api_auth_key

import "github.com/watermint/toolbox/essentials/api/api_auth"

func NewSession(session api_auth.KeySession, repository api_auth.KeyRepository) api_auth.KeySession {
	return &sessionImpl{
		session:    session,
		repository: repository,
	}
}

type sessionImpl struct {
	session    api_auth.KeySession
	repository api_auth.KeyRepository
}

func (z sessionImpl) Start(session api_auth.KeySessionData) (entity api_auth.KeyEntity, err error) {
	entity, found := z.repository.Get(session.AppData.AppKeyName, session.PeerName)
	if found {
		return entity, nil
	}
	entity, err = z.session.Start(session)
	if err != nil {
		return api_auth.KeyEntity{}, err
	}
	z.repository.Put(entity)
	return entity, nil
}
