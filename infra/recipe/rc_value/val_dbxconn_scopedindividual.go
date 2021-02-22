package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
	"strings"
)

func newValueDbxConnScopedIndividual(peerName string) rc_recipe.Value {
	return &ValueDbxConnScopedIndividual{
		conn:     dbx_conn_impl.NewConnScopedIndividual(peerName),
		peerName: peerName,
	}
}

type ValueDbxConnScopedIndividual struct {
	conn     dbx_conn.ConnScopedIndividual
	peerName string
}

func (z *ValueDbxConnScopedIndividual) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*dbx_conn.ConnScopedIndividual)(nil)).Elem()) {
		return newValueDbxConnScopedIndividual(z.peerName)
	}
	return nil
}

func (z *ValueDbxConnScopedIndividual) Bind() interface{} {
	return &z.peerName
}

func (z *ValueDbxConnScopedIndividual) Init() (v interface{}) {
	return z.conn
}

func (z *ValueDbxConnScopedIndividual) ApplyPreset(v0 interface{}) {
	z.conn = v0.(dbx_conn.ConnScopedIndividual)
	z.peerName = z.conn.PeerName()
}

func (z *ValueDbxConnScopedIndividual) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueDbxConnScopedIndividual) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scopes":   strings.Join(z.conn.Scopes(), ","),
		"appType":  z.conn.ScopeLabel(),
	}
}

func (z *ValueDbxConnScopedIndividual) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueDbxConnScopedIndividual) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = dbx_conn_impl.NewConnScopedIndividual(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueDbxConnScopedIndividual) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnScopedIndividual) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnScopedIndividual) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnScopedIndividual) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), z.conn.Scopes()
}

func (z *ValueDbxConnScopedIndividual) ValueText() string {
	return z.peerName
}
