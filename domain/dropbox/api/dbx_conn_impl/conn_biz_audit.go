package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnBusinessAudit(name string) dbx_conn.ConnBusinessAudit {
	cba := &connBusinessAudit{name: name}
	return cba
}

type connBusinessAudit struct {
	name string
	ctx  dbx_context.Context
}

func (z *connBusinessAudit) ServiceName() string {
	return api_conn.ServiceDropboxBusiness
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

func (z *connBusinessAudit) PeerName() string {
	return z.name
}

func (z *connBusinessAudit) Context() dbx_context.Context {
	return z.ctx
}

func (z *connBusinessAudit) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect([]string{z.ScopeLabel()}, z.name, ctl, dbx_auth.NewLegacyApp(ctl))
	return err
}

func (z *connBusinessAudit) IsBusinessAudit() {
}
