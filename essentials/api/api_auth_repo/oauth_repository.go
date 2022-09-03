package api_auth_repo

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewOAuth(repo api_auth2.Repository) api_auth2.OAuthRepository {
	return &oAuthRepository{
		repo: repo,
	}
}

type oAuthRepository struct {
	repo api_auth2.Repository
}

func (z oAuthRepository) Put(entity api_auth2.OAuthEntity) {
	z.repo.Put(entity.Entity())
}

func (z oAuthRepository) Get(keyName string, scopes []string, peerName string) (entity api_auth2.OAuthEntity, found bool) {
	l := esl.Default()
	e, found := z.repo.Get(keyName, api_auth2.OAuthScopeSerialize(scopes), peerName)
	if found {
		if oe, err := api_auth2.DeserializeOAuthEntity(e); err != nil {
			l.Debug("Unable to deserialize", esl.Error(err))
			return entity, false
		} else {
			return oe, true
		}
	}
	return entity, false
}

func (z oAuthRepository) Delete(keyName string, scopes []string, peerName string) {
	z.repo.Delete(keyName, api_auth2.OAuthScopeSerialize(scopes), peerName)
}

func (z oAuthRepository) List(keyName string, scopes []string) (entities []api_auth2.OAuthEntity) {
	l := esl.Default()
	entities = make([]api_auth2.OAuthEntity, 0)
	result := z.repo.List(keyName, api_auth2.OAuthScopeSerialize(scopes))
	for _, e0 := range result {
		e, err := api_auth2.DeserializeOAuthEntity(e0)
		if err != nil {
			l.Debug("unable to deserialize", esl.Error(err))
			continue
		}
		entities = append(entities, e)
	}
	return entities
}

func (z oAuthRepository) Close() {
	z.repo.Close()
}
