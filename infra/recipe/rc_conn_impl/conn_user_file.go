package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

func NewConnUserFile(name string) rc_conn.ConnUserFile {
	return &connUserFile{name: name}
}

type connUserFile struct {
	name string
	ctx  api_context.Context
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

func (z *connUserFile) Name() string {
	return z.name
}

func (z *connUserFile) Context() api_context.Context {
	return z.ctx
}

func (z *connUserFile) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, ctl)
	return err
}

func (z *connUserFile) IsUserFile() {
}
