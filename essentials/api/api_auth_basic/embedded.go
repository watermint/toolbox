package api_auth_basic

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func NewEmbedded(entity api_auth.BasicEntity) api_auth.BasicSession {
	return &embeddedImpl{
		entity: entity,
	}
}

type embeddedImpl struct {
	entity api_auth.BasicEntity
}

func (z embeddedImpl) Start(session api_auth.BasicSessionData) (entity api_auth.BasicEntity, err error) {
	if z.entity.KeyName == session.AppData.AppKeyName &&
		z.entity.PeerName == session.PeerName {
		return z.entity, nil
	}
	return api_auth.BasicEntity{}, app_definitions.ErrorUserCancelled
}
