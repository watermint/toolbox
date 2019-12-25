package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

func NewConnBusinessAudit(name string) rc_conn.ConnBusinessAudit {
	return &connBusinessAudit{name: name}
}

type connBusinessAudit struct {
	name string
	ctx  api_context.Context
}

func (z connBusinessAudit) ScopeLabel() string {
	return api_auth.DropboxTokenBusinessAudit
}

func (z connBusinessAudit) IsPersonal() bool {
	return false
}

func (z connBusinessAudit) IsBusiness() bool {
	return true
}

func (z connBusinessAudit) SetName(name string) {
	z.name = name
}

func (z *connBusinessAudit) Name() string {
	return z.name
}

func (z *connBusinessAudit) Context() api_context.Context {
	return z.ctx
}

func (z *connBusinessAudit) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.ScopeLabel(), z.name, ctl)
	return err
}

func (z *connBusinessAudit) IsBusinessAudit() {
}
