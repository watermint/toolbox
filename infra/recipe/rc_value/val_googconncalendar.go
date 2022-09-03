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

func newValueGoogConnCalendar(peerName string) rc_recipe.Value {
	return &ValueGoogConnCalendar{
		conn:     goog_conn_impl.NewConnGoogleCalendar(peerName),
		peerName: peerName,
	}
}

type ValueGoogConnCalendar struct {
	conn     goog_conn.ConnGoogleCalendar
	peerName string
}

func (z *ValueGoogConnCalendar) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*goog_conn.ConnGoogleCalendar)(nil)).Elem()) {
		return newValueGoogConnCalendar(z.peerName)
	}
	return nil
}

func (z *ValueGoogConnCalendar) Bind() interface{} {
	return &z.peerName
}

func (z *ValueGoogConnCalendar) Init() (v interface{}) {
	return z.conn
}

func (z *ValueGoogConnCalendar) ApplyPreset(v0 interface{}) {
	z.conn = v0.(goog_conn.ConnGoogleCalendar)
	z.peerName = z.conn.PeerName()
}

func (z *ValueGoogConnCalendar) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueGoogConnCalendar) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueGoogConnCalendar) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueGoogConnCalendar) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = goog_conn_impl.NewConnGoogleCalendar(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueGoogConnCalendar) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGoogConnCalendar) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueGoogConnCalendar) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}

func (z *ValueGoogConnCalendar) ValueText() string {
	return z.peerName
}

func (z *ValueGoogConnCalendar) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
