package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

func NewConnBusinessFile(name string) rc_conn.ConnBusinessFile {
	return &connBusinessFile{name: name}
}

type connBusinessFile struct {
	name string
	ctx  api_context.Context
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

func (z *connBusinessFile) Name() string {
	return z.name
}

func (z *connBusinessFile) Context() api_context.Context {
	return z.ctx
}

func (z *connBusinessFile) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, ctl)
	return err
}

func (z *connBusinessFile) IsBusinessFile() {
}
