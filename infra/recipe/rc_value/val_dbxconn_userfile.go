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

func newValueDbxConnUserFile(peerName string) rc_recipe.Value {
	v := &ValueDbxConnUserFile{peerName: peerName}
	v.conn = dbx_conn_impl.NewConnUserFile(peerName)
	return v
}

type ValueDbxConnUserFile struct {
	conn     dbx_conn.ConnUserFile
	peerName string
}

func (z *ValueDbxConnUserFile) Spec() (typeName string, typeAttr interface{}) {
	return ut_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueDbxConnUserFile) ValueText() string {
	return z.peerName
}

func (z *ValueDbxConnUserFile) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*dbx_conn.ConnUserFile)(nil)).Elem()) {
		return newValueDbxConnUserFile(z.peerName)
	}
	return nil
}

func (z *ValueDbxConnUserFile) Bind() interface{} {
	return &z.peerName
}

func (z *ValueDbxConnUserFile) Init() (v interface{}) {
	return z.conn
}

func (z *ValueDbxConnUserFile) ApplyPreset(v0 interface{}) {
	z.conn = v0.(dbx_conn.ConnUserFile)
	z.peerName = z.conn.PeerName()
}

func (z *ValueDbxConnUserFile) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	dbx_conn_impl.EnsurePreVerify(z.conn)
	return z.conn
}

func (z *ValueDbxConnUserFile) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnUserFile) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnUserFile) Conn() (conn dbx_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnUserFile) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
