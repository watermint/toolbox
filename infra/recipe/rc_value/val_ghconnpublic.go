package rc_value

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueGhConnGithubPublic() rc_recipe.Value {
	return &ValueGhConnGithubPublic{
		conn: &gh_conn_impl.ConnGithubPublic{},
	}
}

type ValueGhConnGithubPublic struct {
	conn gh_conn.ConnGithubPublic
}

func (z *ValueGhConnGithubPublic) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*gh_conn.ConnGithubPublic)(nil)).Elem()) {
		return newValueGhConnGithubPublic()
	}
	return nil
}

func (z *ValueGhConnGithubPublic) Bind() interface{} {
	return nil
}

func (z *ValueGhConnGithubPublic) Init() (v interface{}) {
	return z.conn
}

func (z *ValueGhConnGithubPublic) ApplyPreset(v0 interface{}) {
	z.conn = v0.(gh_conn.ConnGithubPublic)
}

func (z *ValueGhConnGithubPublic) Apply() (v interface{}) {
	return z.conn
}

func (z *ValueGhConnGithubPublic) Debug() interface{} {
	return map[string]string{
		"public": "true",
	}
}

func (z *ValueGhConnGithubPublic) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGhConnGithubPublic) SpinDown(ctl app_control.Control) error {
	return nil
}
