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
)

func newValueDbxConnBusinessAudit(peerName string) rc_recipe.Value {
	v := &ValueDbxConnBusinessAudit{peerName: peerName}
	v.conn = dbx_conn_impl.NewConnBusinessAudit(peerName)
	return v
}

type ValueDbxConnBusinessAudit struct {
	conn     dbx_conn.ConnBusinessAudit
	peerName string
}

func (z *ValueDbxConnBusinessAudit) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueDbxConnBusinessAudit) ValueText() string {
	return z.peerName
}

func (z *ValueDbxConnBusinessAudit) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*dbx_conn.ConnBusinessAudit)(nil)).Elem()) {
		return newValueDbxConnBusinessAudit(z.peerName)
	}
	return nil
}

func (z *ValueDbxConnBusinessAudit) Bind() interface{} {
	return &z.peerName
}

func (z *ValueDbxConnBusinessAudit) Init() (v interface{}) {
	return z.conn
}

func (z *ValueDbxConnBusinessAudit) ApplyPreset(v0 interface{}) {
	z.conn = v0.(dbx_conn.ConnBusinessAudit)
	z.peerName = z.conn.PeerName()
}

func (z *ValueDbxConnBusinessAudit) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueDbxConnBusinessAudit) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueDbxConnBusinessAudit) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = dbx_conn_impl.NewConnBusinessAudit(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueDbxConnBusinessAudit) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnBusinessAudit) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnBusinessAudit) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnBusinessAudit) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
