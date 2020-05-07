package gh_conn_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_auth"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type ConnGithubPublic struct {
	name string
	ctl  app_control.Control
}

func (z *ConnGithubPublic) ScopeLabel() string {
	return gh_auth.ScopeNoScope
}

func (z *ConnGithubPublic) PeerName() string {
	return z.name
}

func (z *ConnGithubPublic) SetPeerName(name string) {
	z.name = name
}

func (z *ConnGithubPublic) IsPublic() bool {
	return true
}

func (z *ConnGithubPublic) Connect(ctl app_control.Control) (err error) {
	z.ctl = ctl
	return nil
}

func (z *ConnGithubPublic) Context() gh_context.Context {
	return gh_context_impl.New(z.ctl, nil)
}
