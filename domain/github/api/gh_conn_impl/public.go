package gh_conn_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/api/gh_client_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func NewConnGithubPublic(name string) api_conn.Connection {
	return &ConnGithubPublic{
		name: name,
	}
}

type ConnGithubPublic struct {
	name string
	ctl  app_control.Control
}

func (z *ConnGithubPublic) AppKeyName() string {
	return app_definitions.AppKeyGithubPublic
}

func (z *ConnGithubPublic) ScopeLabel() string {
	return app_definitions.ScopeLabelGithub
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

func (z *ConnGithubPublic) Client() gh_client.Client {
	return gh_client_impl.New("public", z.ctl, api_auth.NewNoAuthOAuthEntity())
}
