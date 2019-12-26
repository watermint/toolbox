package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueFdFileRowFeed(name string) Value {
	v := &ValueFdFileRowFeed{name: name}
	v.rf = fd_file_impl.NewRowFeed(name)
	return v
}

type ValueFdFileRowFeed struct {
	name string
	rf   fd_file.RowFeed
	path string
}

func (z *ValueFdFileRowFeed) ValueText() string {
	return z.path
}

func (z *ValueFdFileRowFeed) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
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

func (z *ValueFdFileRowFeed) ApplyPreset(v0 interface{}) {
	z.rf = v0.(fd_file.RowFeed)
	z.path = z.rf.FilePath()
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

func (z *ValueFdFileRowFeed) Feed() (feed fd_file.RowFeed, valid bool) {
	return z.rf, true
}
