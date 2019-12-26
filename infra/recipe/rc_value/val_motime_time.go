package rc_value

import (
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"reflect"
)

func newValueMoTimeTime(name string) Value {
	v := &ValueMoTimeTime{name: name}
	v.time = mo_time.Zero()
	return v
}

type ValueMoTimeTime struct {
	name     string
	dateTime string
	time     mo_time.Time
}

func (z *ValueMoTimeTime) Accept(t reflect.Type, name string) Value {
	if t.Implements(reflect.TypeOf((*mo_time.Time)(nil)).Elem()) {
		return newValueMoTimeTime(name)
	}
	return nil
}

func (z *ValueMoTimeTime) Bind() interface{} {
	return &z.dateTime
}

func (z *ValueMoTimeTime) Init() (v interface{}) {
	return z.time
}

func (z *ValueMoTimeTime) Apply() (v interface{}) {
	return z.time
}

func (z *ValueMoTimeTime) Debug() interface{} {
	return map[string]string{
		"dateTime": z.dateTime,
		"time":     z.time.String(),
	}
}

func (z *ValueMoTimeTime) SpinUp(ctl app_control.Control) (err error) {
	z.time, err = mo_time.New(z.dateTime)
	if err != nil {
		return err
	}
	return nil
}

func (z *ValueMoTimeTime) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoTimeTime) IsFeed() (feed fd_file.RowFeed, valid bool) {
	return nil, false
}

func (z *ValueMoTimeTime) IsReport() (report rp_model.Report, valid bool) {
	return nil, false
}

func (z *ValueMoTimeTime) IsConn() (conn rc_conn.ConnDropboxApi, valid bool) {
	return nil, false
}
