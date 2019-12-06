package api_request_impl

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	RpcEndpoint     = "api.dropboxapi.com"
	ContentEndpoint = "content.dropboxapi.com"
)

func RpcRequestUrl(endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, endpoint)
}

func ContentRequestUrl(endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", ContentEndpoint, endpoint)
}

func NewPpcRequest(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_request.Request {

	req := &rpcRequestImpl{
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

func NewDownloadRequest(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_request.Request {

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
	content io.Reader,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_request.Request {

	l := app_root.Log()
	upload, err := ioutil.ReadAll(content)
	if err != nil {
		l.Debug("Unable to read", zap.Error(err))
		return &badRequest{err: err}
	}

	req := &uploadRequestImpl{
		ctx:      ctx,
		endpoint: endpoint,
		dbxReq: &dbxRequest{
			asMemberId: asMemberId,
			asAdminId:  asAdminId,
			base:       base,
			token:      token,
		},
		uploadBytes: upload,
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
	token      api_auth.TokenContainer
}

func (z *dbxRequest) decorate(req *http.Request) (r *http.Request, err error) {
	l := app_root.Log()

	if z.token.TokenType != api_auth.DropboxTokenNoAuth {
		req.Header.Add(api_request.ReqHeaderAuthorization, "Bearer "+z.token.Token)
	}
	if z.asAdminId != "" {
		req.Header.Add(api_request.ReqHeaderSelectAdmin, z.asAdminId)
	}
	if z.asMemberId != "" {
		req.Header.Add(api_request.ReqHeaderSelectUser, z.asMemberId)
	}
	if z.base != nil {
		pr, err := json.Marshal(z.base)
		if err != nil {
			l.Debug("unable to marshal path root", zap.Error(err))
			return nil, err
		}
		req.Header.Add(api_request.ReqHeaderPathRoot, string(pr))
	}
	return req, nil
}
