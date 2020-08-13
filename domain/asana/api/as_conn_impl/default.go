package as_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/asana/api/as_auth"
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/domain/asana/api/as_context"
	"github.com/watermint/toolbox/domain/asana/api/as_context_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnAsana(name string) as_conn.ConnAsanaApi {
	return &connAsanaApi{
		name:  name,
		scope: as_auth.ScopeDefault,
	}
}

type connAsanaApi struct {
	name  string
	ctx   as_context.Context
	scope string
}

func (z *connAsanaApi) Connect(ctl app_control.Control) (err error) {
	ac, useMock, err := api_conn_impl.Connect(z.Scopes(), z.name, as_auth.New(ctl), ctl)
	if useMock {
		z.ctx = as_context_impl.NewMock(ctl)
		return nil
	}
	if replay, enabled := ctl.Feature().IsTestWithReplay(); enabled {
		z.ctx = as_context_impl.NewReplayMock(ctl, replay)
		return nil
	}
	if ac != nil {
		z.ctx = as_context_impl.New(ctl, ac)
		return nil
	}
	if err != nil {
		return err
	} else {
		return errors.New("unknown state")
	}
}

func (z *connAsanaApi) PeerName() string {
	return z.name
}

func (z *connAsanaApi) SetPeerName(name string) {
	z.name = name
}

func (z *connAsanaApi) ScopeLabel() string {
	return api_auth.Asana
}

func (z *connAsanaApi) ServiceName() string {
	return api_conn.ServiceAsana
}

func (z *connAsanaApi) SetScopes(scopes ...string) {
	l := z.ctx.Log()
	switch len(z.scope) {
	case 0:
		l.Debug("No scope defined, fallback to default")
		z.scope = as_auth.ScopeDefault
	case 1:
		z.scope = scopes[0]
	default:
		l.Debug("More than one scope defined, fallback to default")
		z.scope = as_auth.ScopeDefault
	}
}

func (z *connAsanaApi) Scopes() []string {
	return []string{z.scope}
}

func (z *connAsanaApi) Context() as_context.Context {
	return z.ctx
}
