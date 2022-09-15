package gh_conn_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_auth"
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/api/gh_client_impl"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnGithubRepo(peerName string) gh_conn.ConnGithubRepo {
	return &ConnGithubRepo{
		peerName: peerName,
	}
}

type ConnGithubRepo struct {
	peerName string
	ctx      gh_client.Client
}

func (z *ConnGithubRepo) ServiceName() string {
	return api_conn.ServiceGithub
}

func (z *ConnGithubRepo) ScopeLabel() string {
	return gh_auth.ScopeLabelRepo
}

func (z *ConnGithubRepo) Connect(ctl app_control.Control) (err error) {
	session := api_auth.OAuthSessionData{
		AppData:  gh_auth.Github,
		PeerName: z.peerName,
		Scopes:   []string{gh_auth.ScopeRepo},
	}
	entity, useMock, err := api_conn_impl.OAuthConnectByRedirect(session, ctl)
	if useMock {
		z.ctx = gh_client_impl.NewMock(z.peerName, ctl)
		return nil
	}
	if err != nil {
		return err
	}
	z.ctx = gh_client_impl.New(z.peerName, ctl, entity)
	return nil
}

func (z *ConnGithubRepo) PeerName() string {
	return z.peerName
}

func (z *ConnGithubRepo) SetPeerName(name string) {
	z.peerName = name
}

func (z *ConnGithubRepo) Client() gh_client.Client {
	return z.ctx
}

func (z *ConnGithubRepo) IsRepo() bool {
	return true
}
