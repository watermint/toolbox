package rc_value

import (
	"github.com/watermint/toolbox/domain/slack/api/work_conn"
	"github.com/watermint/toolbox/domain/slack/api/work_conn_impl"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueSlack(peerName string) rc_recipe.Value {
	return &ValueSlackConn{
		conn:     work_conn_impl.NewSlackApi(peerName),
		peerName: peerName,
	}
}

type ValueSlackConn struct {
	conn     work_conn.ConnSlackApi
	peerName string
}

func (z *ValueSlackConn) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*work_conn.ConnSlackApi)(nil)).Elem()) {
		return newValueSlack(name)
	}
	return nil
}

func (z *ValueSlackConn) Bind() interface{} {
	return &z.peerName
}

func (z *ValueSlackConn) Init() (v interface{}) {
	return z.conn
}

func (z *ValueSlackConn) ApplyPreset(v0 interface{}) {
	z.conn = v0.(work_conn.ConnSlackApi)
	z.peerName = z.conn.PeerName()
}

func (z *ValueSlackConn) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueSlackConn) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueSlackConn) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueSlackConn) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueSlackConn) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}

func (z *ValueSlackConn) ValueText() string {
	return z.peerName
}

func (z *ValueSlackConn) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
