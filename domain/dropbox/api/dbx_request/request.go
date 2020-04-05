package dbx_request

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"net/http"
)

const (
	RpcEndpoint     = "api.dropboxapi.com"
	NotifyEndpoint  = "notify.dropboxapi.com"
	ContentEndpoint = "content.dropboxapi.com"
)

func RpcRequestUrl(base, endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", base, endpoint)
}

func ContentRequestUrl(endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", ContentEndpoint, endpoint)
}

func NewPpcRequest(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.Context,
	endpointBase string) api_request.Request {

	req := &rpcRequestImpl{
		ctx:          ctx,
		endpoint:     endpoint,
		endpointBase: endpointBase,
		dbxReq: &dbxRequest{
			asMemberId: asMemberId,
			asAdminId:  asAdminId,
			base:       base,
			token:      token,
		},
	}
	return req
}

func NewDownloadRequest(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.Context) api_request.Request {

	req := &downloadRequestImpl{
		ctx:      ctx,
		endpoint: endpoint,
		dbxReq: &dbxRequest{
			asMemberId: asMemberId,
			asAdminId:  asAdminId,
			base:       base,
			token:      token,
		},
	}
	return req
}

func NewUploadRequest(ctx api_context.Context,
	endpoint string,
	content ut_io.ReadRewinder,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.Context) api_request.Request {

	req := &uploadRequestImpl{
		ctx:      ctx,
		endpoint: endpoint,
		dbxReq: &dbxRequest{
			asMemberId: asMemberId,
			asAdminId:  asAdminId,
			base:       base,
			token:      token,
		},
		content: content,
	}
	return req
}

type badRequest struct {
	err error
}

func (z *badRequest) ContentLength() int64 {
	return 0
}

func (z *badRequest) ParamString() string {
	return ""
}

func (z *badRequest) Param(p interface{}) api_request.Request {
	return &badRequest{
		err: z.err,
	}
}

func (z *badRequest) Call() (res api_response.Response, err error) {
	return nil, z.err
}

func (z *badRequest) Endpoint() string {
	return ""
}

func (z *badRequest) Url() string {
	return ""
}

func (z *badRequest) Headers() map[string]string {
	return map[string]string{}
}

func (z *badRequest) Method() string {
	return ""
}

func (z *badRequest) Make() (req *http.Request, err error) {
	return nil, z.err
}

type dbxRequest struct {
	asMemberId string
	asAdminId  string
	base       api_context.PathRoot
	token      api_auth.Context
}

func (z *dbxRequest) decorate(req *http.Request) (r *http.Request, err error) {
	l := app_root.Log()

	req.Header.Add(api_request.ReqHeaderUserAgent, app.UserAgent())
	if !z.token.IsNoAuth() {
		req.Header.Add(api_request.ReqHeaderAuthorization, "Bearer "+z.token.Token().AccessToken)
	}
	if z.asAdminId != "" {
		req.Header.Add(api_request.ReqHeaderDropboxApiSelectAdmin, z.asAdminId)
	}
	if z.asMemberId != "" {
		req.Header.Add(api_request.ReqHeaderDropboxApiSelectUser, z.asMemberId)
	}
	if z.base != nil {
		pr, err := json.Marshal(z.base)
		if err != nil {
			l.Debug("unable to marshal path root", zap.Error(err))
			return nil, err
		}
		req.Header.Add(api_request.ReqHeaderDropboxApiPathRoot, string(pr))
	}
	return req, nil
}
