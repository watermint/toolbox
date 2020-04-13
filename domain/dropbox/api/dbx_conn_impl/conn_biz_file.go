package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnBusinessFile(name string) dbx_conn.ConnBusinessFile {
	cbf := &connBusinessFile{name: name}
	return cbf
}

type connBusinessFile struct {
	name   string
	verify bool
	ctx    dbx_context.Context
}

func (z *connBusinessFile) SetPreVerify(enabled bool) {
	z.verify = enabled
}

func (z *connBusinessFile) IsPreVerify() bool {
	return z.verify
}

func (z *connBusinessFile) ScopeLabel() string {
	return api_auth.DropboxTokenBusinessFile
}

func (z *connBusinessFile) IsPersonal() bool {
	return false
}

func (z *connBusinessFile) IsBusiness() bool {
	return true
}

func (z *connBusinessFile) SetPeerName(name string) {
	z.name = name
}

func (z *connBusinessFile) PeerName() string {
	return z.name
}

func (z *connBusinessFile) Context() dbx_context.Context {
	return z.ctx
}

func (z *connBusinessFile) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, z.verify, ctl)
	return err
}

func (z *connBusinessFile) IsBusinessFile() {
}
