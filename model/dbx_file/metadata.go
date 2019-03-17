package dbx_file

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Metadata struct {
	Path                            string `json:"path"`
	IncludeMediaInfo                bool   `json:"include_media_info"`
	IncludeDeleted                  bool   `json:"include_deleted"`
	IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members"`

	AsMemberId string                      `json:"-"`
	AsAdminId  string                      `json:"-"`
	PathRoot   interface{}                 `json:"-"`
	OnError    func(err error) bool        `json:"-"`
	OnFolder   func(folder *Folder) bool   `json:"-"`
	OnFile     func(file *File) bool       `json:"-"`
	OnDelete   func(deleted *Deleted) bool `json:"-"`
}

func (z *Metadata) Get(c *dbx_api.DbxContext) bool {
	ep := EntryParser{
		Logger:   c.Log(),
		OnError:  z.OnError,
		OnFile:   z.OnFile,
		OnFolder: z.OnFolder,
		OnDelete: z.OnDelete,
	}
	req := dbx_rpc.RpcRequest{
		AsMemberId: z.AsMemberId,
		AsAdminId:  z.AsAdminId,
		PathRoot:   z.PathRoot,
		Endpoint:   "files/get_metadata",
		Param:      z,
	}
	res, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}
	m := gjson.Parse(res.Body)
	if !m.Exists() {
		c.Log().Debug("response body is not a JSON", zap.String("response", res.Body))
		return false
	}
	return ep.Parse(m)
}
