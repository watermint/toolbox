package api_context_impl

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/domain/infra/api_async"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_list_impl"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/domain/infra/api_rpc_impl"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func New(ec *app.ExecContext, token api_auth.Token, options ...api_context.Option) api_context.Context {
	c := &contextImpl{
		ec:     ec,
		dt:     token,
		client: &http.Client{},
	}
	for _, op := range options {
		op(c)
	}
	return c
}

func NewContextNoAuth(ec *app.ExecContext) api_context.Context {
	DefaultClientTimeout := time.Duration(60) * time.Second
	c := contextImpl{
		ec:     ec,
		dt:     nil,
		noAuth: true,
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
	}
	return &c
}

type contextImpl struct {
	ec         *app.ExecContext
	dt         api_auth.Token
	noAuth     bool
	client     *http.Client
	asMemberId string
	asAdminId  string
	basePath   api_context.Base
}

func (z *contextImpl) ClientTimeout(second int) {
	z.client.Timeout = time.Duration(second) * time.Second
}

func (z *contextImpl) AsMemberId(teamMemberId string) {
	z.asMemberId = teamMemberId
}

func (z *contextImpl) AsAdminId(teamMemberId string) {
	z.asAdminId = teamMemberId
}

func (z *contextImpl) BasePath(pathRoot api_context.Base) {
	z.basePath = pathRoot
}

func (z *contextImpl) Log() *zap.Logger {
	return z.ec.Log()
}

func (z *contextImpl) Msg(key string) app_ui.UIMessage {
	return z.ec.Msg(key)
}

func (z *contextImpl) Request(endpoint string) api_rpc.Request {
	return api_rpc_impl.New(z.ec, endpoint, z.asMemberId, z.asAdminId, z.basePath, z.dt)
}

func (z *contextImpl) List(endpoint string) api_list.List {
	return api_list_impl.New(z.ec, z, endpoint, z.asMemberId, z.asAdminId, z.basePath, z.dt)
}

func (z *contextImpl) Async(endpoint string) api_async.Async {
	panic("implement me")
}
