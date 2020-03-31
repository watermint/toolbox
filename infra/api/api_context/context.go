package api_context

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"net/http"
)

var (
	ErrorIncompatibleContextType = errors.New("incompatible context type")
)

type Context interface {
	Log() *zap.Logger

	Rpc(endpoint string) api_request.Request
	Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request
	Download(endpoint string) api_request.Request

	NoRetryOnError() Context
	IsNoRetry() bool
	Hash() string
	NoAuth() Context
	MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error)
}

type DropboxApiContext interface {
	Context
	Notify(endpoint string) api_request.Request
	List(endpoint string) api_list.List
	Async(endpoint string) api_async.Async
	AsMemberId(teamMemberId string) DropboxApiContext
	AsAdminId(teamMemberId string) DropboxApiContext
	WithPath(pathRoot PathRoot) DropboxApiContext
}

type CaptureContext interface {
	Capture() *zap.Logger
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
