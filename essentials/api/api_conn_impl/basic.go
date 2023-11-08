package api_conn_impl

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_basic"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func BasicConnect(session api_auth.BasicSessionData, ctl app_control.Control, opts ...api_auth_basic.ConsoleOpt) (entity api_auth.BasicEntity, useMock bool, err error) {
	if isTest, mock, err := IsTestMode(ctl); isTest {
		return api_auth.NewNoAuthBasicEntity(), mock, err
	}
	s := api_auth_basic.NewSession(
		api_auth_basic.NewConsole(ctl, opts...),
		ctl.AuthRepository(),
	)
	entity, err = s.Start(session)
	return entity, false, err
}
