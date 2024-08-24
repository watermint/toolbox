package fg_conn_impl

import (
	"github.com/watermint/toolbox/domain/figma/api/fg_auth"
	"github.com/watermint/toolbox/domain/figma/api/fg_client"
	"github.com/watermint/toolbox/domain/figma/api/fg_client_impl"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func NewConnFigma(name string) fg_conn.ConnFigmaApi {
	return &connFigmaApi{
		peerName: name,
		scopes:   []string{fg_auth.ScopeFileRead},
	}
}

type connFigmaApi struct {
	peerName string
	client   fg_client.Client
	scopes   []string
}

func (z *connFigmaApi) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connFigmaApi) Scopes() []string {
	return z.scopes
}

func (z *connFigmaApi) Connect(ctl app_control.Control) (err error) {
	session := api_auth.OAuthSessionData{
		AppData:  fg_auth.Figma,
		PeerName: z.peerName,
		Scopes:   z.Scopes(),
	}
	entity, useMock, err := api_conn_impl.OAuthConnectByRedirect(session, ctl)
	if useMock {
		z.client = fg_client_impl.NewMock(z.peerName, ctl)
		return nil
	}
	if err != nil {
		return err
	}
	z.client = fg_client_impl.New(z.peerName, ctl, entity)
	return nil
}

func (z *connFigmaApi) PeerName() string {
	return z.peerName
}

func (z *connFigmaApi) SetPeerName(name string) {
	z.peerName = name
}

func (z *connFigmaApi) ScopeLabel() string {
	return app_definitions.ScopeLabelFigma
}

func (z *connFigmaApi) AppKeyName() string {
	return app_definitions.AppKeyFigma
}

func (z *connFigmaApi) Client() fg_client.Client {
	return z.client
}
