package rc_value

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_conn_impl"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
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

func (z *ValueGhConnGithubRepo) Conn() (conn api_conn.Connection, valid bool) {
	return z.conn, true
}

func (z *ValueGhConnGithubRepo) ValueText() string {
	return z.peerName
}

func (z *ValueGhConnGithubRepo) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.conn), nil
}

func (z *ValueGhConnGithubRepo) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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

func (z *ValueGhConnGithubRepo) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.peerName, nil
}

func (z *ValueGhConnGithubRepo) Restore(v es_json.Json, ctl app_control.Control) error {
	if peerName, found := v.String(); found {
		z.conn = gh_conn_impl.NewConnGithubRepo(peerName)
		z.peerName = peerName
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueGhConnGithubRepo) SpinUp(ctl app_control.Control) error {
	return z.conn.Connect(ctl)
}

func (z *ValueGhConnGithubRepo) SpinDown(ctl app_control.Control) error {
	return nil
}
