package api_auth_repo

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewOAuth(repo api_auth.Repository) api_auth.OAuthRepository {
	return &oAuthRepository{
		repo: repo,
	}
}

type oAuthRepository struct {
	repo api_auth.Repository
}

func (z oAuthRepository) Put(entity api_auth.OAuthEntity) {
	z.repo.Put(entity.Entity())
}

func (z oAuthRepository) Get(keyName string, scopes []string, peerName string) (entity api_auth.OAuthEntity, found bool) {
	l := esl.Default()
	e, found := z.repo.Get(keyName, api_auth.OAuthScopeSerialize(scopes), peerName)
	if found {
		if oe, err := api_auth.DeserializeOAuthEntity(e); err != nil {
			l.Debug("Unable to deserialize", esl.Error(err))
			return api_auth.OAuthEntity{}, false
		} else {
			return oe, true
		}
	}
	return api_auth.OAuthEntity{}, false
}

func (z oAuthRepository) Delete(keyName string, scopes []string, peerName string) {
	z.repo.Delete(keyName, api_auth.OAuthScopeSerialize(scopes), peerName)
}

func (z oAuthRepository) List(keyName string, scopes []string) (entities []api_auth.OAuthEntity) {
	l := esl.Default()
	entities = make([]api_auth.OAuthEntity, 0)
	result := z.repo.List(keyName, api_auth.OAuthScopeSerialize(scopes))
	for _, e0 := range result {
		e, err := api_auth.DeserializeOAuthEntity(e0)
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
