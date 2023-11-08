package api_conn_impl

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_key"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func KeyConnect(data api_auth.KeySessionData, ctl app_control.Control, askMsg app_msg.Message) (entity api_auth.KeyEntity, useMock bool, err error) {
	if isTest, mock, err := IsTestMode(ctl); isTest {
		return api_auth.NewNoAuthKeyEntity(), mock, err
	}
	s := api_auth_key.NewSession(
		api_auth_key.NewConsole(ctl, askMsg),
		api_auth_repo.NewKey(ctl.AuthRepository()),
	)
	entity, err = s.Start(data)
	return entity, false, err
}
