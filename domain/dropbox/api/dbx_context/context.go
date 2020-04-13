package dbx_context

import (
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Context interface {
	api_context.Context
	api_context.PostContext
	api_context.AsyncContext
	api_context.ListContext
	api_context.UploadContext
	api_context.DownloadContext

	AsMemberId(teamMemberId string) Context
	AsAdminId(teamMemberId string) Context
	WithPath(pathRoot PathRoot) Context
}

type NoAuthContext interface {
	api_context.Context
	api_context.NotifyContext
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
