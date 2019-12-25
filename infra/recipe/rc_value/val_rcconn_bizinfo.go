package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueRcConnBusinessInfo(name string) Value {
	v := &ValueRcConnBusinessInfo{name: name}
	v.conn = rc_conn_impl.NewConnBusinessInfo(name)
	return v
}

type ValueRcConnBusinessInfo struct {
	conn rc_conn.ConnBusinessInfo
	name string
}

func (z *ValueRcConnBusinessInfo) Accept(t reflect.Type, name string) Value {
	if t.Implements(reflect.TypeOf((*rc_conn.ConnBusinessInfo)(nil)).Elem()) {
		return newValueRpModelRowReport(name)
	}
	return nil
}

func (z *ValueRcConnBusinessInfo) Bind() interface{} {
	return &z.name
}

func (z *ValueRcConnBusinessInfo) Init() (v interface{}) {
	return z.conn
}

func (z *ValueRcConnBusinessInfo) Apply() (v interface{}) {
	z.conn.SetName(z.name)
	return z.conn
}

func (z *ValueRcConnBusinessInfo) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueRcConnBusinessInfo) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueRcConnBusinessInfo) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueRcConnBusinessInfo) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueRcConnBusinessInfo) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return z.conn, true
}

func (z *ValueRcConnBusinessInfo) Debug() interface{} {
	return map[string]string{
		"peerName": z.name,
		"scope":    z.conn.ScopeLabel(),
	}
}
