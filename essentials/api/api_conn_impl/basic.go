package api_conn_impl

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_basic"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func BasicConnect(session api_auth.BasicSessionData, ctl app_control.Control) (entity api_auth.BasicEntity, useMock bool, err error) {
	if isTest, mock, err := isTestMode(ctl); isTest {
		return api_auth.NewNoAuthBasicEntity(), mock, err
	}
	s := api_auth_basic.NewRepository(
		api_auth_basic.NewConsole(ctl),
		ctl.AuthRepository(),
	)
	entity, err = s.Start(session)
	return entity, false, err
}
