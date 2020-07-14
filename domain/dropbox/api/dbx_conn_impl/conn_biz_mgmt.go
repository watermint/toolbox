package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnBusinessMgmt(name string) dbx_conn.ConnBusinessMgmt {
	cbm := &connBusinessMgmt{name: name}
	return cbm
}

type connBusinessMgmt struct {
	name string
	ctx  dbx_context.Context
}

func (z *connBusinessMgmt) ServiceName() string {
	return api_conn.ServiceDropboxBusiness
}

func (z *connBusinessMgmt) ScopeLabel() string {
	return api_auth.DropboxTokenBusinessManagement
}

func (z *connBusinessMgmt) IsPersonal() bool {
	return false
}

func (z *connBusinessMgmt) IsBusiness() bool {
	return true
}

func (z *connBusinessMgmt) SetPeerName(name string) {
	z.name = name
}

func (z *connBusinessMgmt) PeerName() string {
	return z.name
}

func (z *connBusinessMgmt) Context() dbx_context.Context {
	return z.ctx
}

func (z *connBusinessMgmt) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect([]string{z.ScopeLabel()}, z.name, ctl, dbx_auth.NewLegacyApp(ctl))
	return err
}

func (z *connBusinessMgmt) IsBusinessMgmt() {
}
