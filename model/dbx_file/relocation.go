package dbx_file

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Relocation struct {
	FromPath               string `json:"from_path"`
	ToPath                 string `json:"to_path"`
	AllowSharedFolder      bool   `json:"allow_shared_folder"`
	AllowOwnershipTransfer bool   `json:"allow_ownership_transfer"`
	AutoRename             bool   `json:"autorename"`

	AsMemberId string                      `json:"-"`
	OnError    func(err error) bool        `json:"-"`
	OnFolder   func(folder *Folder) bool   `json:"-"`
	OnFile     func(file *File) bool       `json:"-"`
	OnDelete   func(deleted *Deleted) bool `json:"-"`
}

func (z *Relocation) relocation(c *dbx_api.DbxContext, endpoint string) bool {
	ep := EntryParser{
		Logger:   c.Log().With(zap.String("endpoint", endpoint)),
		OnError:  z.OnError,
		OnFile:   z.OnFile,
		OnFolder: z.OnFolder,
		OnDelete: z.OnDelete,
	}
	req := dbx_rpc.RpcRequest{
		AsMemberId: z.AsMemberId,
		Endpoint:   endpoint,
		Param:      z,
	}
	res, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}
	m := gjson.Get(res.Body, "metadata")
	if !m.Exists() {
		c.Log().Debug("response `metadata` not found", zap.String("response", res.Body))
		return false
	}
	return ep.Parse(m)
}

func (z *Relocation) Copy(c *dbx_api.DbxContext) bool {
	return z.relocation(c, "files/copy_v2")
}

func (z *Relocation) Move(c *dbx_api.DbxContext) bool {
	return z.relocation(c, "files/move_v2")
}

type MassRelocation struct {
	Force      bool
	AsMemberId string
	OnError    func(err error) bool
}

func (z *MassRelocation) Copy(c *dbx_api.DbxContext, src, dest string) bool {
	return z.relocation(c, src, dest, "files/copy_v2")
}

func (z *MassRelocation) Move(c *dbx_api.DbxContext, src, dest string) bool {
	return z.relocation(c, src, dest, "files/move_v2")
}

func (z *MassRelocation) relocation(c *dbx_api.DbxContext, src, dest, endpoint string) bool {
	panic("implement me!")
}
