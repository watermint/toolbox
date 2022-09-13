package api_auth_repo

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewBasic(repo api_auth.Repository) api_auth.BasicRepository {
	return &basicRepoImpl{
		repo: repo,
	}
}

type basicRepoImpl struct {
	repo api_auth.Repository
}

func (z basicRepoImpl) Put(entity api_auth.BasicEntity) {
	z.repo.Put(entity.Entity())
}

func (z basicRepoImpl) Get(keyName string, peerName string) (entity api_auth.BasicEntity, found bool) {
	l := esl.Default()
	e, found := z.repo.Get(keyName, "", peerName)
	if found {
		if be, err := api_auth.DeserializeBasicEntity(e); err != nil {
			l.Debug("Unable to deserialize", esl.Error(err))
			return api_auth.BasicEntity{}, false
		} else {
			return be, true
		}
	}
	return api_auth.BasicEntity{}, false
}

func (z basicRepoImpl) Delete(keyName string, peerName string) {
	z.repo.Delete(keyName, "", peerName)
}

func (z basicRepoImpl) List(keyName string) (entities []api_auth.BasicEntity) {
	l := esl.Default()
	entities = make([]api_auth.BasicEntity, 0)
	result := z.repo.List(keyName, "")
	for _, e0 := range result {
		e, err := api_auth.DeserializeBasicEntity(e0)
		if err != nil {
			l.Debug("unable to deserialize", esl.Error(err))
			continue
		}
		entities = append(entities, e)
	}
	return entities
}

func (z basicRepoImpl) Close() {
	z.repo.Close()
}
