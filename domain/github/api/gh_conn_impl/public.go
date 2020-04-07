package gh_conn_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type ConnGithubPublic struct {
	ctl app_control.Control
}

func (z *ConnGithubPublic) IsPublic() bool {
	return true
}

func (z *ConnGithubPublic) Connect(ctl app_control.Control) (err error) {
	z.ctl = ctl
	return nil
}

func (z *ConnGithubPublic) Context() gh_context.Context {
	return gh_context_impl.NewNoAuth(z.ctl)
}
