package app_conn_impl

import (
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/experimental/app_kitchen"
)

type ConnBusinessMgmt struct {
	PeerName string
}

func (z *ConnBusinessMgmt) Context(kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	c := api_auth_impl.NewKc(kitchen, api_auth_impl.PeerName(z.PeerName))
	ctx, err = c.Auth(api_auth.DropboxTokenBusinessManagement)
	return
}

func (*ConnBusinessMgmt) IsBusinessMgmt() {
}
