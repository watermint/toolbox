package rc_value

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueGhConnGithubRepo(peerName string) rc_recipe.Value {
	return &ValueGhConnGithubRepo{
		peerName: peerName,
		conn:     gh_conn_impl.NewConnGithubRepo(peerName),
	}
}

type ValueGhConnGithubRepo struct {
	conn     gh_conn.ConnGithubRepo
	peerName string
}

func (z *ValueGhConnGithubRepo) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*gh_conn.ConnGithubRepo)(nil)).Elem()) {
		return newValueGhConnGithubRepo(z.peerName)
	}
	return nil
}

func (z *ValueGhConnGithubRepo) Bind() interface{} {
	return &z.peerName
}

func (z *ValueGhConnGithubRepo) Init() (v interface{}) {
	return z.conn
}

func (z *ValueGhConnGithubRepo) ApplyPreset(v0 interface{}) {
	z.conn = v0.(gh_conn.ConnGithubRepo)
	z.peerName = z.conn.PeerName()
}

func (z *ValueGhConnGithubRepo) Apply() (v interface{}) {
	z.conn.SetPeerName(z.peerName)
	return z.conn
}

func (z *ValueGhConnGithubRepo) Debug() interface{} {
	return map[string]string{
		"peerName": z.peerName,
		"scope":    z.conn.ScopeLabel(),
	}
}

func (z *ValueGhConnGithubRepo) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGhConnGithubRepo) SpinDown(ctl app_control.Control) error {
	return nil
}
