package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueRcConnBusinessAudit(peerName string) Value {
	v := &ValueRcConnBusinessAudit{peerName: peerName}
	v.conn = rc_conn_impl.NewConnBusinessAudit(peerName)
	return v
}

type ValueRcConnBusinessAudit struct {
	conn     rc_conn.ConnBusinessAudit
	peerName string
}

func (z *ValueRcConnBusinessAudit) ValueText() string {
	return z.peerName
}

func (z *ValueRcConnBusinessAudit) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Implements(reflect.TypeOf((*rc_conn.ConnBusinessAudit)(nil)).Elem()) {
		return newValueRcConnBusinessAudit(z.peerName)
	}
	return nil
}

func (z *ValueRcConnBusinessAudit) Bind() interface{} {
	return &z.peerName
}

func (z *ValueRcConnBusinessAudit) Init() (v interface{}) {
	return z.conn
}

func (z *ValueRcConnBusinessAudit) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueRcConnBusinessAudit) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnBusinessAudit) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnBusinessAudit) Conn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnBusinessAudit) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
