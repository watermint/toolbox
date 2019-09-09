package sq_test

import (
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_test"
	app2 "github.com/watermint/toolbox/legacy/app"
)

func DoTestTeamTask(test func(biz service.Business)) {
	peerName := api_test.TestPeerName
	ec := app2.NewExecContextForTest()
	defer ec.Shutdown()
	if !api_auth_impl.IsCacheAvailable(ec, peerName) {
		return
	}

	ctxMgmt, err := api_auth_impl.Auth(ec, api_auth_impl.PeerName(peerName), api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}
	ctxFile, err := api_auth_impl.Auth(ec, api_auth_impl.PeerName(peerName), api_auth_impl.BusinessFile())
	if err != nil {
		return
	}
	biz, err := service.New(ctxMgmt, ctxFile)
	if err != nil {
		return
	}
	test(biz)
}
