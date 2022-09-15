package rc_value

import (
	"github.com/watermint/toolbox/domain/hellosign/api/hs_conn"
	"github.com/watermint/toolbox/domain/hellosign/api/hs_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueHsConn(peerName string) rc_recipe.Value {
	return &ValueHsConn{
		peerName: peerName,
		conn:     hs_conn_impl.NewConnHelloSign(peerName),
	}
}

type ValueHsConn struct {
	conn     hs_conn.ConnHelloSignApi
	peerName string
}

func (z *ValueHsConn) ValueText() string {
	return z.peerName
}

func (z *ValueHsConn) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueHsConn) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*hs_conn.ConnHelloSignApi)(nil)).Elem()) {
		return newValueHsConn(z.peerName)
	}
	return nil
}

func (z *ValueHsConn) Bind() interface{} {
	return &z.peerName
}

func (z *ValueHsConn) Init() (v interface{}) {
	return z.conn
}

func (z *ValueHsConn) ApplyPreset(v0 interface{}) {
	z.conn = v0.(hs_conn.ConnHelloSignApi)
	z.peerName = z.conn.PeerName()
}

func (z *ValueHsConn) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueHsConn) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
	}
}

func (z *ValueHsConn) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueHsConn) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = hs_conn_impl.NewConnHelloSign(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueHsConn) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueHsConn) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueHsConn) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}
