package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueRcConnBusinessMgmt(peerName string) rc_recipe.Value {
	v := &ValueRcConnBusinessMgmt{peerName: peerName}
	v.conn = rc_conn_impl.NewConnBusinessMgmt(peerName)
	return v
}

type ValueRcConnBusinessMgmt struct {
	conn     rc_conn.ConnBusinessMgmt
	peerName string
}

func (z *ValueRcConnBusinessMgmt) ValueText() string {
	return z.peerName
}

func (z *ValueRcConnBusinessMgmt) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*rc_conn.ConnBusinessMgmt)(nil)).Elem()) {
		return newValueRcConnBusinessMgmt(z.peerName)
	}
	return nil
}

func (z *ValueRcConnBusinessMgmt) Bind() interface{} {
	return &z.peerName
}

func (z *ValueRcConnBusinessMgmt) Init() (v interface{}) {
	return z.conn
}

func (z *ValueRcConnBusinessMgmt) ApplyPreset(v0 interface{}) {
	z.conn = v0.(rc_conn.ConnBusinessMgmt)
	z.peerName = z.conn.Name()
}

func (z *ValueRcConnBusinessMgmt) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	rc_conn_impl.EnsurePreVerify(z.conn)
	return z.conn
}

func (z *ValueRcConnBusinessMgmt) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnBusinessMgmt) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnBusinessMgmt) Conn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnBusinessMgmt) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
