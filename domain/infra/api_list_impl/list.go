package api_list_impl

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/domain/infra/api_rpc_impl"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

func New(ec *app.ExecContext, ctx api_context.Context, endpoint string, asMemberId, asAdminId string, base api_context.Base, token api_auth.Token) api_list.List {
	return &listImpl{
		ec:              ec,
		ctx:             ctx,
		requestEndpoint: endpoint,
		asMemberId:      asMemberId,
		asAdminId:       asAdminId,
		base:            base,
	}
}

type listImpl struct {
	ctx              api_context.Context
	ec               *app.ExecContext
	asMemberId       string
	asAdminId        string
	base             api_context.Base
	param            interface{}
	token            api_auth.Token
	useHasMore       bool
	resultTag        string
	requestEndpoint  string
	continueEndpoint string
	onEntry          func(res api_list.ListEntry) error
	onResponse       func(res api_rpc.Response) error
	onFailure        func(err error) error
}

func (z *listImpl) Param(param interface{}) api_list.List {
	z.param = param
	return z
}

func (z *listImpl) Continue(endpoint string) api_list.List {
	z.continueEndpoint = endpoint
	return z
}

func (z *listImpl) UseHasMore(use bool) api_list.List {
	z.useHasMore = use
	return z
}

func (z *listImpl) ResultTag(tag string) api_list.List {
	z.resultTag = tag
	return z
}

func (z *listImpl) OnFailure(failure func(err error) error) api_list.List {
	z.onFailure = failure
	return z
}

func (z *listImpl) OnResponse(response func(res api_rpc.Response) error) api_list.List {
	z.onResponse = response
	return z
}

func (z *listImpl) OnEntry(entry func(entry api_list.ListEntry) error) api_list.List {
	z.onEntry = entry
	return z
}

func (z *listImpl) Call() (err error) {
	ls := dbx_rpc.RpcList{
		EndpointList:         z.requestEndpoint,
		EndpointListContinue: z.continueEndpoint,
		//PathRoot: z.base, //TODO: require type change
		UseHasMore: z.useHasMore,
		AsMemberId: z.asMemberId,
		AsAdminId:  z.asAdminId,
		ResultTag:  z.resultTag,
		OnError: func(err error) bool {
			z.ctx.Log().Debug("error", zap.Error(err))
			if z.onFailure != nil {
				return z.onFailure(err) != nil
			}
			return false
		},
		OnResponse: func(res *dbx_rpc.RpcResponse) bool {
			if z.onResponse == nil {
				return true
			}

			if z.onResponse(api_rpc_impl.NewSuccessResponse(res)) != nil {
				return false
			}
			return true
		},
		OnEntry: func(result gjson.Result) bool {
			if z.onEntry != nil {
				e := &listEntryImpl{
					entry: result,
				}
				if err := z.onEntry(e); err != nil {
					z.ctx.Log().Debug("onEntry returned error", zap.Error(err))
					return false
				}
				// fall through
			}
			return true
		},
	}

	rpcReq := z.ctx.Request(z.requestEndpoint).Param(z.param)
	switch qi := rpcReq.(type) {
	case *api_rpc_impl.RequestImpl:
		if !ls.List(qi.DbxApiContext(), z.param) {
			return errors.New("operation failed")
		}
		return nil
	}
	panic("invalid state")
}

type listEntryImpl struct {
	entry gjson.Result
}

func (z *listEntryImpl) Json() (res gjson.Result, err error) {
	return z.entry, nil
}

func (z *listEntryImpl) Model(v interface{}) error {
	return api_parser.ParseModel(v, z.entry)
}

func (z *listEntryImpl) ModelWithPath(v interface{}, path string) error {
	return api_parser.ParseModel(v, z.entry.Get(path))
}
