package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDbxConnBusinessInfo(peerName string) rc_recipe.Value {
	v := &ValueDbxConnBusinessInfo{peerName: peerName}
	v.conn = dbx_conn_impl.NewConnBusinessInfo(peerName)
	return v
}

type ValueDbxConnBusinessInfo struct {
	conn     dbx_conn.ConnBusinessInfo
	peerName string
}

func (z *ValueDbxConnBusinessInfo) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueDbxConnBusinessInfo) ValueText() string {
	return z.peerName
}

func (z *ValueDbxConnBusinessInfo) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*dbx_conn.ConnBusinessInfo)(nil)).Elem()) {
		return newValueDbxConnBusinessInfo(z.peerName)
	}
	return nil
}

func (z *ValueDbxConnBusinessInfo) Bind() interface{} {
	return &z.peerName
}

func (z *ValueDbxConnBusinessInfo) Init() (v interface{}) {
	return z.conn
}

func (z *ValueDbxConnBusinessInfo) ApplyPreset(v0 interface{}) {
	z.conn = v0.(dbx_conn.ConnBusinessInfo)
	z.peerName = z.conn.PeerName()
}

func (z *ValueDbxConnBusinessInfo) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	dbx_conn_impl.EnsurePreVerify(z.conn)
	return z.conn
}

func (z *ValueDbxConnBusinessInfo) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnBusinessInfo) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnBusinessInfo) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnBusinessInfo) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
