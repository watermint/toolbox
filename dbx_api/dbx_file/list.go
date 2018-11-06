package dbx_file

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"go.uber.org/zap"
)

type ListFolder struct {
	AsMemberId                      string
	AsAdminId                       string
	IncludeMediaInfo                bool
	IncludeDeleted                  bool
	IncludeHasExplicitSharedMembers bool
	IncludeMountedFolders           bool

	OnError  func(annotation dbx_api.ErrorAnnotation) bool
	OnFolder func(folder *Folder) bool
	OnFile   func(file *File) bool
	OnDelete func(deleted *Deleted) bool
}

func (l *ListFolder) List(c *dbx_api.Context, path string) bool {
	type ListParam struct {
		Path                            string `json:"path"`
		IncludeMediaInfo                bool   `json:"include_media_info,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
		IncludeMountedFolders           bool   `json:"include_mounted_folders,omitempty"`
	}
	lp := ListParam{
		Path:                            path,
		IncludeMediaInfo:                l.IncludeMediaInfo,
		IncludeDeleted:                  l.IncludeDeleted,
		IncludeHasExplicitSharedMembers: l.IncludeHasExplicitSharedMembers,
		IncludeMountedFolders:           l.IncludeMountedFolders,
	}
	ep := EntryParser{
		Logger:   c.Log().With(zap.String("operation", "files/list_folder")),
		OnError:  l.OnError,
		OnFile:   l.OnFile,
		OnFolder: l.OnFolder,
		OnDelete: l.OnDelete,
	}

	list := dbx_rpc.RpcList{
		AsMemberId:           l.AsMemberId,
		AsAdminId:            l.AsAdminId,
		EndpointList:         "files/list_folder",
		EndpointListContinue: "files/list_folder/continue",
		UseHasMore:           true,
		ResultTag:            "entries",
		OnError:              l.OnError,
		OnEntry: func(result gjson.Result) bool {
			return ep.Parse(result)
		},
	}
	return list.List(c, lp)
}
