package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
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
	return ut_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueDbxConnBusinessAudit) ValueText() string {
	return z.peerName
}

func (z *ValueDbxConnBusinessAudit) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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
	dbx_conn_impl.EnsurePreVerify(z.conn)
	return z.conn
}

func (z *ValueDbxConnBusinessAudit) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnBusinessAudit) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnBusinessAudit) Conn() (conn dbx_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnBusinessAudit) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
