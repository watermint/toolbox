package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
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

func (z *ValueRcConnBusinessAudit) Fork(ctl app_control.Control) Value {
	v := &ValueRcConnBusinessAudit{}
	v.peerName = z.peerName
	v.conn = rc_conn_impl.NewConnBusinessAudit(z.peerName)
	v.conn.SetName(z.peerName)
	return v
}

func (z *ValueRcConnBusinessAudit) Accept(t reflect.Type, name string) Value {
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
	z.conn.SetName(z.peerName)
	return z.conn
}

func (z *ValueRcConnBusinessAudit) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnBusinessAudit) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnBusinessAudit) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueRcConnBusinessAudit) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueRcConnBusinessAudit) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnBusinessAudit) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}
