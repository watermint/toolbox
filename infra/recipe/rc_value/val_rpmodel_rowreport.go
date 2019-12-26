package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"reflect"
)

func newValueRpModelRowReport(name string) Value {
	v := &ValueRpModelRowReport{name: name}
	v.rep = rp_model_impl.NewRowReport(name)
	return v
}

type ValueRpModelRowReport struct {
	name string
	rep  *rp_model_impl.RowReport
}

func (z *ValueRpModelRowReport) Accept(t reflect.Type, name string) Value {
	if t.Implements(reflect.TypeOf((*rp_model.RowReport)(nil)).Elem()) {
		return newValueRpModelRowReport(name)
	}
	return nil
}

func (z *ValueRpModelRowReport) Bind() interface{} {
	return nil
}

func (z *ValueRpModelRowReport) Init() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelRowReport) Apply() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelRowReport) SpinUp(ctl app_control.Control) error {
	// Report will not automatically open
	z.rep.SetCtl(ctl)
	return nil
}

func (z *ValueRpModelRowReport) SpinDown(ctl app_control.Control) error {
	z.rep.Close()
	return nil
}

func (z *ValueRpModelRowReport) Debug() interface{} {
	return nil
}

func (z *ValueRpModelRowReport) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueRpModelRowReport) IsReport() (report rp_model.Report, valid bool) {
	return z.rep, true
}

func (z *ValueRpModelRowReport) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}
