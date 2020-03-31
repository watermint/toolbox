package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

func NewConnBusinessAudit(name string) rc_conn.ConnBusinessAudit {
	cba := &connBusinessAudit{name: name}
	return cba
}

type connBusinessAudit struct {
	name   string
	verify bool
	ctx    api_context.DropboxApiContext
}

func (z *connBusinessAudit) SetPreVerify(enabled bool) {
	z.verify = enabled
}

func (z *connBusinessAudit) IsPreVerify() bool {
	return z.verify
}

func (z *connBusinessAudit) ScopeLabel() string {
	return api_auth.DropboxTokenBusinessAudit
}

func (z *connBusinessAudit) IsPersonal() bool {
	return false
}

func (z *connBusinessAudit) IsBusiness() bool {
	return true
}

func (z *connBusinessAudit) SetPeerName(name string) {
	z.name = name
}

func (z *connBusinessAudit) Name() string {
	return z.name
}

func (z *connBusinessAudit) Context() api_context.DropboxApiContext {
	return z.ctx
}

func (z *connBusinessAudit) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, z.verify, ctl)
	return err
}

func (z *connBusinessAudit) IsBusinessAudit() {
}
