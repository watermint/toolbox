package gh_conn_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_auth"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnGithubRepo(name string) gh_conn.ConnGithubRepo {
	return &ConnGithubRepo{
		name: name,
	}
}

type ConnGithubRepo struct {
	name string
	ctx  gh_context.Context
}

func (z *ConnGithubRepo) ServiceName() string {
	return api_conn.ServiceGithub
}

func (z *ConnGithubRepo) ScopeLabel() string {
	return gh_auth.ScopeLabelRepo
}

func (z *ConnGithubRepo) Connect(ctl app_control.Control) (err error) {
	ac, useMock, err := api_conn_impl.Connect([]string{gh_auth.ScopeRepo}, z.name, gh_auth.NewApp(ctl), ctl)
	if useMock {
		z.ctx = gh_context_impl.NewMock(ctl)
		return nil
	}
	if ac != nil {
		z.ctx = gh_context_impl.New(ctl, ac)
		return nil
	}
	return err
}

func (z *ConnGithubRepo) PeerName() string {
	return z.name
}

func (z *ConnGithubRepo) SetPeerName(name string) {
	z.name = name
}

func (z *ConnGithubRepo) Context() gh_context.Context {
	return z.ctx
}

func (z *ConnGithubRepo) IsRepo() bool {
	return true
}
