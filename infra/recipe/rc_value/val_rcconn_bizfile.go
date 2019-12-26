package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueRcConnBusinessFile(peerName string) Value {
	v := &ValueRcConnBusinessFile{peerName: peerName}
	v.conn = rc_conn_impl.NewConnBusinessFile(peerName)
	return v
}

type ValueRcConnBusinessFile struct {
	conn     rc_conn.ConnBusinessFile
	peerName string
}

func (z *ValueRcConnBusinessFile) Fork(ctl app_control.Control) Value {
	v := &ValueRcConnBusinessFile{}
	v.peerName = z.peerName
	v.conn = rc_conn_impl.NewConnBusinessFile(z.peerName)
	v.conn.SetName(z.peerName)
	return v
}

func (z *ValueRcConnBusinessFile) Accept(t reflect.Type, name string) Value {
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

func (z *ValueRcConnBusinessFile) Apply() (v interface{}) {
	z.conn.SetName(z.peerName)
	return z.conn
}

func (z *ValueRcConnBusinessFile) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnBusinessFile) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnBusinessFile) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueRcConnBusinessFile) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueRcConnBusinessFile) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnBusinessFile) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
