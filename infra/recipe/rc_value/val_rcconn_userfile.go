package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
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

func (z *ValueRcConnUserFile) Fork(ctl app_control.Control) Value {
	v := &ValueRcConnUserFile{}
	v.peerName = z.peerName
	v.conn = rc_conn_impl.NewConnUserFile(z.peerName)
	v.conn.SetName(z.peerName)
	return v
}

func (z *ValueRcConnUserFile) Accept(t reflect.Type, name string) Value {
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

func (z *ValueRcConnUserFile) Apply() (v interface{}) {
	z.conn.SetName(z.peerName)
	return z.conn
}

func (z *ValueRcConnUserFile) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnUserFile) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnUserFile) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueRcConnUserFile) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueRcConnUserFile) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnUserFile) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
