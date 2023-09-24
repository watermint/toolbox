package rc_value

import (
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueGoogConnTranslate(peerName string) rc_recipe.Value {
	return &ValueGoogConnTranslate{
		conn:     goog_conn_impl.NewConnGoogleTranslate(peerName),
		peerName: peerName,
	}
}

type ValueGoogConnTranslate struct {
	conn     goog_conn.ConnGoogleTranslate
	peerName string
}

func (z *ValueGoogConnTranslate) ValueText() string {
	return z.peerName
}

func (z *ValueGoogConnTranslate) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueGoogConnTranslate) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*goog_conn.ConnGoogleTranslate)(nil)).Elem()) {
		return newValueGoogConnTranslate(z.peerName)
	}
	return nil

}

func (z *ValueGoogConnTranslate) Bind() interface{} {
	return &z.peerName
}

func (z *ValueGoogConnTranslate) Init() (v interface{}) {
	return z.conn
}

func (z *ValueGoogConnTranslate) ApplyPreset(v0 interface{}) {
	z.conn = v0.(goog_conn.ConnGoogleTranslate)
	z.peerName = z.conn.PeerName()
}

func (z *ValueGoogConnTranslate) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn

}

func (z *ValueGoogConnTranslate) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueGoogConnTranslate) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueGoogConnTranslate) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = goog_conn_impl.NewConnGoogleTranslate(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueGoogConnTranslate) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGoogConnTranslate) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueGoogConnTranslate) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}
