package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"reflect"
)

func newValueRpModelTransactionReport(name string) Value {
	v := &ValueRpModelTransactionReport{name: name}
	v.rep = rp_model_impl.NewTransactionReport(name)
	return v
}

type ValueRpModelTransactionReport struct {
	name string
	rep  *rp_model_impl.TransactionReport
}

func (z *ValueRpModelTransactionReport) Accept(t reflect.Type, name string) Value {
	if t.Implements(reflect.TypeOf((*rp_model.TransactionReport)(nil)).Elem()) {
		return newValueRpModelTransactionReport(name)
	}
	return nil
}

func (z *ValueRpModelTransactionReport) Bind() interface{} {
	return nil
}

func (z *ValueRpModelTransactionReport) Init() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelTransactionReport) Apply() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelTransactionReport) SpinUp(ctl app_control.Control) error {
	// Report will not automatically open
	z.rep.SetCtl(ctl)
	return nil
}

func (z *ValueRpModelTransactionReport) SpinDown(ctl app_control.Control) error {
	z.rep.Close()
	return nil
}

func (z *ValueRpModelTransactionReport) Debug() interface{} {
	return nil
}

func (z *ValueRpModelTransactionReport) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueRpModelTransactionReport) IsReport() (report rp_model.Report, valid bool) {
	return z.rep, true
}

func (z *ValueRpModelTransactionReport) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}
