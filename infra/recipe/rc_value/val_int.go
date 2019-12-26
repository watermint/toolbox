package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueInt() Value {
	return &ValueInt{}
}

type ValueInt struct {
	v int
}

func (z *ValueInt) Accept(t reflect.Type, name string) Value {
	if t.Kind() == reflect.Int {
		return newValueInt()
	}
	return nil
}

func (z *ValueInt) Fork(ctl app_control.Control) Value {
	v := &ValueInt{}
	v.v = z.v
	return v
}

func (z *ValueInt) Bind() interface{} {
	return &z.v
}

func (z *ValueInt) Init() (v interface{}) {
	return z.v
}

func (z *ValueInt) Apply() (v interface{}) {
	return z.v
}

func (z *ValueInt) Debug() interface{} {
	return z.v
}

func (z *ValueInt) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueInt) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueInt) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueInt) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueInt) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}
