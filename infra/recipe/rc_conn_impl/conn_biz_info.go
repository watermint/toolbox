package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

func NewConnBusinessInfo(name string) rc_conn.ConnBusinessInfo {
	return &connBusinessInfo{name: name}
}

type connBusinessInfo struct {
	name string
	ctx  api_context.Context
}

func (z *connBusinessInfo) ScopeLabel() string {
	return api_auth.DropboxTokenBusinessInfo
}

func (z *connBusinessInfo) IsPersonal() bool {
	return false
}

func (z *connBusinessInfo) IsBusiness() bool {
	return true
}

func (z *connBusinessInfo) SetPeerName(name string) {
	z.name = name
}

func (z *connBusinessInfo) Name() string {
	return z.name
}

func (z *connBusinessInfo) Context() api_context.Context {
	return z.ctx
}

func (z *connBusinessInfo) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, ctl)
	return err
}

func (z *connBusinessInfo) IsBusinessInfo() {
}
