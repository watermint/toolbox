package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDbxConnBusinessMgmt(peerName string) rc_recipe.Value {
	v := &ValueDbxConnBusinessMgmt{peerName: peerName}
	v.conn = dbx_conn_impl.NewConnBusinessMgmt(peerName)
	return v
}

type ValueDbxConnBusinessMgmt struct {
	conn     dbx_conn.ConnBusinessMgmt
	peerName string
}

func (z *ValueDbxConnBusinessMgmt) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueDbxConnBusinessMgmt) ValueText() string {
	return z.peerName
}

func (z *ValueDbxConnBusinessMgmt) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*dbx_conn.ConnBusinessMgmt)(nil)).Elem()) {
		return newValueDbxConnBusinessMgmt(z.peerName)
	}
	return nil
}

func (z *ValueDbxConnBusinessMgmt) Bind() interface{} {
	return &z.peerName
}

func (z *ValueDbxConnBusinessMgmt) Init() (v interface{}) {
	return z.conn
}

func (z *ValueDbxConnBusinessMgmt) ApplyPreset(v0 interface{}) {
	z.conn = v0.(dbx_conn.ConnBusinessMgmt)
	z.peerName = z.conn.PeerName()
}

func (z *ValueDbxConnBusinessMgmt) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	dbx_conn_impl.EnsurePreVerify(z.conn)
	return z.conn
}

func (z *ValueDbxConnBusinessMgmt) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnBusinessMgmt) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnBusinessMgmt) Conn() (conn dbx_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnBusinessMgmt) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
