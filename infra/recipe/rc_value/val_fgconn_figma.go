package rc_value

import (
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueFgConnFigmaFileRead(peerName string) rc_recipe.Value {
	return &ValueFgConnFigmaFileRead{
		peerName: peerName,
		conn:     fg_conn_impl.NewConnFigma(peerName),
	}
}

type ValueFgConnFigmaFileRead struct {
	conn     fg_conn.ConnFigmaApi
	peerName string
}

func (z *ValueFgConnFigmaFileRead) ValueText() string {
	return z.peerName
}

func (z *ValueFgConnFigmaFileRead) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*fg_conn.ConnFigmaApi)(nil)).Elem()) {
		return newValueFgConnFigmaFileRead(z.peerName)
	}
	return nil
}

func (z *ValueFgConnFigmaFileRead) Bind() interface{} {
	return &z.peerName
}

func (z *ValueFgConnFigmaFileRead) Init() (v interface{}) {
	return z.conn
}

func (z *ValueFgConnFigmaFileRead) ApplyPreset(v0 interface{}) {
	z.conn = v0.(fg_conn.ConnFigmaApi)
	z.peerName = z.conn.PeerName()
}

func (z *ValueFgConnFigmaFileRead) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueFgConnFigmaFileRead) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueFgConnFigmaFileRead) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueFgConnFigmaFileRead) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = fg_conn_impl.NewConnFigma(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueFgConnFigmaFileRead) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueFgConnFigmaFileRead) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueFgConnFigmaFileRead) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueFgConnFigmaFileRead) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}
