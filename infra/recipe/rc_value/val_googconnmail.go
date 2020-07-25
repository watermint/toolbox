package rc_value

import (
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_conn_impl"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueGoogConnMail(peerName string) rc_recipe.Value {
	return &ValueGoogConnMail{
		peerName: peerName,
		conn:     goog_conn_impl.NewConnGoogleMail(peerName),
	}
}

type ValueGoogConnMail struct {
	conn     goog_conn.ConnGoogleApi
	peerName string
}

func (z *ValueGoogConnMail) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*goog_conn.ConnGoogleMail)(nil)).Elem()) {
		return newValueGoogConnMail(z.peerName)
	}
	return nil
}

func (z *ValueGoogConnMail) Bind() interface{} {
	return &z.peerName
}

func (z *ValueGoogConnMail) Init() (v interface{}) {
	return z.conn
}

func (z *ValueGoogConnMail) ApplyPreset(v0 interface{}) {
	z.conn = v0.(goog_conn.ConnGoogleMail)
	z.peerName = z.conn.PeerName()
}

func (z *ValueGoogConnMail) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueGoogConnMail) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueGoogConnMail) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGoogConnMail) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueGoogConnMail) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}

func (z *ValueGoogConnMail) ValueText() string {
	return z.peerName
}

func (z *ValueGoogConnMail) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
