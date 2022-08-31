package as_conn_impl

import (
	"github.com/watermint/toolbox/domain/asana/api/as_auth"
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/domain/asana/api/as_client_impl"
	"github.com/watermint/toolbox/domain/asana/api/as_conn"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnAsana(name string) as_conn.ConnAsanaApi {
	return &connAsanaApi{
		peerName: name,
		scope:    as_auth.ScopeDefault,
	}
}

type connAsanaApi struct {
	peerName string
	ctx      as_client.Client
	scope    string
}

func (z *connAsanaApi) Connect(ctl app_control.Control) (err error) {
	session := api_auth.OAuthSessionData{
		AppData:  as_auth.Asana,
		PeerName: z.peerName,
		Scopes:   z.Scopes(),
	}
	entity, useMock, err := api_conn_impl.ConnectByRedirect(session, ctl)
	if useMock {
		z.ctx = as_client_impl.NewMock(z.peerName, ctl)
		return nil
	}
	if replay, enabled := ctl.Feature().IsTestWithSeqReplay(); enabled {
		z.ctx = as_client_impl.NewReplayMock(z.peerName, ctl, replay)
		return nil
	}
	if err != nil {
		return err
	}
	z.ctx = as_client_impl.New(z.peerName, ctl, entity)
	return nil
}

func (z *connAsanaApi) PeerName() string {
	return z.peerName
}

func (z *connAsanaApi) SetPeerName(name string) {
	z.peerName = name
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

func (z *connAsanaApi) Context() as_client.Client {
	return z.ctx
}
