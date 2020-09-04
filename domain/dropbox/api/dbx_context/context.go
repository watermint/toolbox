package dbx_context

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
)

const (
	DropboxApiResHeaderRetryAfter    = "Retry-After"
	DropboxApiResHeaderResult        = "Dropbox-API-Result"
	DropboxApiErrorBadInputParam     = 400
	DropboxApiErrorBadOrExpiredToken = 401
	DropboxApiErrorAccessError       = 403
	DropboxApiErrorEndpointSpecific  = 409
	DropboxApiErrorNoPermission      = 422
	DropboxApiErrorRateLimit         = 429
)

type ListContext interface {
	List(endpoint string) dbx_list.List
}

type AsyncContext interface {
	Async(endpoint string) dbx_async.Async
}

type Context interface {
	api_context.Context
	api_context.QualityContext
	api_context.UI

	Async(endpoint string, d ...api_request.RequestDatum) dbx_async.Async
	List(endpoint string, d ...api_request.RequestDatum) dbx_list.List
	Post(endpoint string, d ...api_request.RequestDatum) dbx_response.Response
	Upload(endpoint string, d ...api_request.RequestDatum) dbx_response.Response
	Download(endpoint string, d ...api_request.RequestDatum) dbx_response.Response
	Notify(endpoint string, d ...api_request.RequestDatum) dbx_response.Response

	AsMemberId(teamMemberId string) Context
	AsAdminId(teamMemberId string) Context
	WithPath(pathRoot PathRoot) Context
	NoAuth() Context
	NoRetry() Context
}

type PathRoot interface {
	Header() string
}

func Home() PathRoot {
	return &homePathRoot{Tag: "home"}
}
func Root(namespaceId string) PathRoot {
	return &rootPathRoot{Tag: "root", Root: namespaceId}
}
func Namespace(namespaceId string) PathRoot {
	return &namespacePathRoot{Tag: "namespace_id", NamespaceId: namespaceId}
}

type homePathRoot struct {
	Tag string `json:".tag"`
}

func (*homePathRoot) Header() string {
	return "{\".tag\":\"home\"}"
}

type rootPathRoot struct {
	Tag  string `json:".tag"`
	Root string `json:"root"`
}

func (z *rootPathRoot) Header() string {
	return "{\".tag\":\"root\",\"root\":\"" + z.Root + "\"}"
}

type namespacePathRoot struct {
	Tag         string `json:".tag"`
	NamespaceId string `json:"namespace_id"`
}

func (z *namespacePathRoot) Header() string {
	return "{\".tag\":\"namespace_id\",\"namespace_id\":\"" + z.NamespaceId + "\"}"
}

func ContentResponseData(res es_response.Response) es_json.Json {
	if j, err := es_json.ParseString(res.Header(DropboxApiResHeaderResult)); err != nil {
		return es_json.Null()
	} else {
		return j
	}
}
