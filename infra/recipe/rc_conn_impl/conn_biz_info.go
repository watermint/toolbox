package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

func NewConnBusinessInfo(name string) rc_conn.ConnBusinessInfo {
	cbi := &connBusinessInfo{name: name}
	return cbi
}

type connBusinessInfo struct {
	name   string
	verify bool
	ctx    api_context.DropboxApiContext
}

func (z *connBusinessInfo) SetPreVerify(enabled bool) {
	z.verify = enabled
}

func (z *connBusinessInfo) IsPreVerify() bool {
	return z.verify
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

func (z *connBusinessInfo) Context() api_context.DropboxApiContext {
	return z.ctx
}

func (z *connBusinessInfo) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, z.verify, ctl)
	return err
}

func (z *connBusinessInfo) IsBusinessInfo() {
}
