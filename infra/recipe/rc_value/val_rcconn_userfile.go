package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueRcConnUserFile(peerName string) Value {
	v := &ValueRcConnUserFile{peerName: peerName}
	v.conn = rc_conn_impl.NewConnUserFile(peerName)
	return v
}

type ValueRcConnUserFile struct {
	conn     rc_conn.ConnUserFile
	peerName string
}

func (z *ValueRcConnUserFile) ValueText() string {
	return z.peerName
}

func (z *ValueRcConnUserFile) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Implements(reflect.TypeOf((*rc_conn.ConnUserFile)(nil)).Elem()) {
		return newValueRcConnUserFile(z.peerName)
	}
	return nil
}

func (z *ValueRcConnUserFile) Bind() interface{} {
	return &z.peerName
}

func (z *ValueRcConnUserFile) Init() (v interface{}) {
	return z.conn
}

func (z *ValueRcConnUserFile) Apply(v0 interface{}) (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueRcConnUserFile) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnUserFile) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnUserFile) Conn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnUserFile) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}