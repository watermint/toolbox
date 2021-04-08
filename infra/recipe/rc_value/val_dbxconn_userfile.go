package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_error_handler"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
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

func (z *ValueDbxConnUserFile) ErrorHandler() rc_error_handler.ErrorHandler {
	return dbx_error.NewHandler()
}

func (z *ValueDbxConnUserFile) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueDbxConnUserFile) ValueText() string {
	return z.peerName
}

func (z *ValueDbxConnUserFile) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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
	return z.conn
}

func (z *ValueDbxConnUserFile) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueDbxConnUserFile) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = dbx_conn_impl.NewConnUserFile(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueDbxConnUserFile) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueDbxConnUserFile) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDbxConnUserFile) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueDbxConnUserFile) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
