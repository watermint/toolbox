package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueFdFileRowFeed(name string) Value {
	v := &ValueFdFileRowFeed{}
	v.rf = fd_file_impl.NewRowFeed(name)
	return v
}

type ValueFdFileRowFeed struct {
	rf   fd_file.RowFeed
	path string
}

func (z *ValueFdFileRowFeed) Accept(t reflect.Type, name string) Value {
	if t.Implements(reflect.TypeOf((*fd_file.RowFeed)(nil)).Elem()) {
		return newValueFdFileRowFeed(name)
	}
	return nil
}

func (z *ValueFdFileRowFeed) Bind() interface{} {
	return &z.path
}

func (z *ValueFdFileRowFeed) Init() (v interface{}) {
	return z.rf
}

func (z *ValueFdFileRowFeed) Apply() (v interface{}) {
	z.rf.SetFilePath(z.path)
	return z.rf
}

func (z *ValueFdFileRowFeed) Debug() interface{} {
	return map[string]string{
		"path": z.path,
	}
}

func (z *ValueFdFileRowFeed) SpinUp(ctl app_control.Control) error {
	return z.rf.Open(ctl)
}

func (z *ValueFdFileRowFeed) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueFdFileRowFeed) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return z.rf, true
}

func (z *ValueFdFileRowFeed) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueFdFileRowFeed) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}
