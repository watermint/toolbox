package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueRcConnUserFile(peerName string) rc_recipe.Value {
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

func (z *ValueRcConnUserFile) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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

func (z *ValueRcConnUserFile) ApplyPreset(v0 interface{}) {
	z.conn = v0.(rc_conn.ConnUserFile)
	z.peerName = z.conn.Name()
}

func (z *ValueRcConnUserFile) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	rc_conn_impl.EnsurePreVerify(z.conn)
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
