package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueBool() Value {
	return &ValueBool{}
}

type ValueBool struct {
	v bool
}

func (z *ValueBool) Fork(ctl app_control.Control) Value {
	v := newValueBool().(*ValueBool)
	v.v = z.v
	return v
}

func (z *ValueBool) Accept(t reflect.Type, name string) Value {
	if t.Kind() == reflect.Bool {
		return newValueBool()
	}
	return nil
}

func (z *ValueBool) Bind() interface{} {
	return &z.v
}

func (z *ValueBool) Init() (v interface{}) {
	return z.v
}

func (z *ValueBool) Apply() (v interface{}) {
	return z.v
}

func (z *ValueBool) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueBool) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueBool) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueBool) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueBool) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}

func (z *ValueBool) Debug() interface{} {
	return z.v
}
