package rc_value

import (
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/domain/asana/api/as_conn_impl"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueAsConnAsana(peerName string) rc_recipe.Value {
	return &ValueAsConnAsana{
		peerName: peerName,
		conn:     as_conn_impl.NewConnAsana(peerName),
	}
}

type ValueAsConnAsana struct {
	conn     as_conn.ConnAsanaApi
	peerName string
}

func (z *ValueAsConnAsana) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*as_conn.ConnAsanaApi)(nil)).Elem()) {
		return newValueAsConnAsana(z.peerName)
	}
	return nil
}

func (z *ValueAsConnAsana) Bind() interface{} {
	return &z.peerName
}

func (z *ValueAsConnAsana) Init() (v interface{}) {
	return z.conn
}

func (z *ValueAsConnAsana) ApplyPreset(v0 interface{}) {
	z.conn = v0.(as_conn.ConnAsanaApi)
	z.peerName = z.conn.PeerName()
}

func (z *ValueAsConnAsana) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueAsConnAsana) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueAsConnAsana) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueAsConnAsana) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = as_conn_impl.NewConnAsana(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueAsConnAsana) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueAsConnAsana) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueAsConnAsana) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}

func (z *ValueAsConnAsana) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
