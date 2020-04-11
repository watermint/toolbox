package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnUserFile(name string) dbx_conn.ConnUserFile {
	cuf := &connUserFile{name: name}
	return cuf
}

type connUserFile struct {
	name   string
	verify bool
	ctx    dbx_context.Context
}

func (z *connUserFile) SetPreVerify(enabled bool) {
	z.verify = enabled
}

func (z *connUserFile) IsPreVerify() bool {
	return z.verify
}

func (z *connUserFile) SetPeerName(name string) {
	z.name = name
}

func (z *connUserFile) ScopeLabel() string {
	return api_auth.DropboxTokenFull
}

func (z *connUserFile) IsPersonal() bool {
	return true
}

func (z *connUserFile) IsBusiness() bool {
	return false
}

func (z *connUserFile) PeerName() string {
	return z.name
}

func (z *connUserFile) Context() dbx_context.Context {
	return z.ctx
}

func (z *connUserFile) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, z.verify, ctl)
	return err
}

func (z *connUserFile) IsUserFile() {
}
