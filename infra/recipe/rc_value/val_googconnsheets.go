package rc_value

import (
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_conn_impl"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueGoogConnSheets(peerName string) rc_recipe.Value {
	return &ValueGoogConnSheets{
		conn:     goog_conn_impl.NewConnGoogleSheets(peerName),
		peerName: peerName,
	}
}

type ValueGoogConnSheets struct {
	conn     goog_conn.ConnGoogleSheets
	peerName string
}

func (z *ValueGoogConnSheets) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*goog_conn.ConnGoogleSheets)(nil)).Elem()) {
		return newValueGoogConnSheets(z.peerName)
	}
	return nil
}

func (z *ValueGoogConnSheets) Bind() interface{} {
	return &z.peerName
}

func (z *ValueGoogConnSheets) Init() (v interface{}) {
	return z.conn
}

func (z *ValueGoogConnSheets) ApplyPreset(v0 interface{}) {
	z.conn = v0.(goog_conn.ConnGoogleSheets)
	z.peerName = z.conn.PeerName()
}

func (z *ValueGoogConnSheets) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueGoogConnSheets) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueGoogConnSheets) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueGoogConnSheets) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = goog_conn_impl.NewConnGoogleSheets(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueGoogConnSheets) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGoogConnSheets) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueGoogConnSheets) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}

func (z *ValueGoogConnSheets) ValueTxt() string {
	return z.peerName
}

func (z *ValueGoogConnSheets) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
