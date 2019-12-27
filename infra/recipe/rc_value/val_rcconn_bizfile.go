package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueRcConnBusinessFile(peerName string) rc_recipe.Value {
	v := &ValueRcConnBusinessFile{peerName: peerName}
	v.conn = rc_conn_impl.NewConnBusinessFile(peerName)
	return v
}

type ValueRcConnBusinessFile struct {
	conn     rc_conn.ConnBusinessFile
	peerName string
}

func (z *ValueRcConnBusinessFile) ValueText() string {
	return z.peerName
}

func (z *ValueRcConnBusinessFile) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*rc_conn.ConnBusinessFile)(nil)).Elem()) {
		return newValueRcConnBusinessFile(z.peerName)
	}
	return nil
}

func (z *ValueRcConnBusinessFile) Bind() interface{} {
	return &z.peerName
}

func (z *ValueRcConnBusinessFile) Init() (v interface{}) {
	return z.conn
}

func (z *ValueRcConnBusinessFile) ApplyPreset(v0 interface{}) {
	z.conn = v0.(rc_conn.ConnBusinessFile)
	z.peerName = z.conn.Name()
}

func (z *ValueRcConnBusinessFile) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueRcConnBusinessFile) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnBusinessFile) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnBusinessFile) Conn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnBusinessFile) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
