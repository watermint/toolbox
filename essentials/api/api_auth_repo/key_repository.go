package api_auth_repo

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewKey(repo api_auth.Repository) api_auth.KeyRepository {
	return &keyRepoImpl{
		repository: repo,
	}
}

type keyRepoImpl struct {
	repository api_auth.Repository
}

func (z keyRepoImpl) Put(entity api_auth.KeyEntity) {
	z.repository.Put(entity.Entity())
}

func (z keyRepoImpl) Get(keyName, peerName string) (entity api_auth.KeyEntity, found bool) {
	l := esl.Default()
	e, found := z.repository.Get(keyName, "", peerName)
	if found {
		return DeserializeKeyEntity(e), true
	}
	l.Debug("Not found", esl.String("keyName", keyName), esl.String("peerName", peerName))
	return api_auth.KeyEntity{}, false
}

func (z keyRepoImpl) Delete(keyName, peerName string) {
	z.repository.Delete(keyName, "", peerName)
}

func (z keyRepoImpl) List(keyName string) (entities []api_auth.KeyEntity) {
	entities = make([]api_auth.KeyEntity, 0)
	result := z.repository.List(keyName, "")
	for _, e0 := range result {
		entities = append(entities, DeserializeKeyEntity(e0))
	}
	return entities
}

func (z keyRepoImpl) Close() {
	z.repository.Close()
}

func DeserializeKeyEntity(entity api_auth.Entity) api_auth.KeyEntity {
	return api_auth.KeyEntity{
		KeyName:  entity.KeyName,
		PeerName: entity.PeerName,
		Credential: api_auth.KeyCredential{
			Key: entity.Credential,
		},
		Description: entity.Description,
		Timestamp:   entity.Timestamp,
	}
}
