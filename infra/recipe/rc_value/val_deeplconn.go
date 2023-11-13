package rc_value

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn"
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDeeplConn(peerName string) rc_recipe.Value {
	return &ValueDeeplConn{
		peerName: peerName,
		conn:     deepl_conn_impl.NewConnDeeplApi(peerName),
	}
}

type ValueDeeplConn struct {
	conn     deepl_conn.ConnDeeplApi
	peerName string
}

func (z *ValueDeeplConn) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueDeeplConn) ValueText() string {
	return z.peerName
}

func (z *ValueDeeplConn) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*deepl_conn.ConnDeeplApi)(nil)).Elem()) {
		return newValueDeeplConn(z.peerName)
	}
	return nil
}

func (z *ValueDeeplConn) Bind() interface{} {
	return &z.peerName
}

func (z *ValueDeeplConn) Init() (v interface{}) {
	return z.conn
}

func (z *ValueDeeplConn) ApplyPreset(v0 interface{}) {
	z.conn = v0.(deepl_conn.ConnDeeplApi)
	z.peerName = z.conn.PeerName()
}

func (z *ValueDeeplConn) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueDeeplConn) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueDeeplConn) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueDeeplConn) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = deepl_conn_impl.NewConnDeeplApi(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueDeeplConn) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDeeplConn) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDeeplConn) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), ""
}
