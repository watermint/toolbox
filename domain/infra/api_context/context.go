package api_context

import (
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/domain/infra/api_async"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"go.uber.org/zap"
)

type Context interface {
	Log() *zap.Logger
	Msg(key string) app_ui.UIMessage
	ErrorMsg(err error) app_ui.UIMessage

	Request(endpoint string) api_rpc.Request
	List(endpoint string) api_list.List
	Async(endpoint string) api_async.Async

	AsMemberId(teamMemberId string) Context
	AsAdminId(teamMemberId string) Context
	BasePath(pathRoot Base) Context
}

type Base interface {
	Value() string
}
type Root interface {
	Base
}
type Namespace interface {
	Base
	NamespaceId() string
}
type Home interface {
	Base
}
