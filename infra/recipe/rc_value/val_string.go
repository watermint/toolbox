package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueString() Value {
	return &ValueString{}
}

type ValueString struct {
	v string
}

func (z *ValueString) Accept(t reflect.Type, name string) Value {
	if t.Kind() == reflect.String {
		return newValueString()
	}
	return nil
}

func (z *ValueString) Fork(ctl app_control.Control) Value {
	v := &ValueString{}
	v.v = z.v
	return v
}

func (z *ValueString) Bind() interface{} {
	return &z.v
}

func (z *ValueString) Init() (v interface{}) {
	return z.v
}

func (z *ValueString) Apply() (v interface{}) {
	return z.v
}

func (z *ValueString) Debug() interface{} {
	return z.v
}

func (z *ValueString) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueString) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueString) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueString) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueString) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}
