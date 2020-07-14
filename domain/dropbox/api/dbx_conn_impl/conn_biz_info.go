package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnBusinessInfo(name string) dbx_conn.ConnBusinessInfo {
	cbi := &connBusinessInfo{name: name}
	return cbi
}

type connBusinessInfo struct {
	name string
	ctx  dbx_context.Context
}

func (z *connBusinessInfo) ServiceName() string {
	return api_conn.ServiceDropboxBusiness
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

func (z *connBusinessInfo) PeerName() string {
	return z.name
}

func (z *connBusinessInfo) Context() dbx_context.Context {
	return z.ctx
}

func (z *connBusinessInfo) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect([]string{z.ScopeLabel()}, z.name, ctl, dbx_auth.NewLegacyApp(ctl))
	return err
}

func (z *connBusinessInfo) IsBusinessInfo() {
}
