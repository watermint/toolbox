package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/api/goog_client_impl"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func NewConnGoogleTranslate(name string) goog_conn.ConnGoogleTranslate {
	return &connTranslateImpl{
		name:   name,
		scopes: []string{},
	}
}

type connTranslateImpl struct {
	name   string
	scopes []string
	ctx    goog_client.Client
}

func (z *connTranslateImpl) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(goog_auth.Sheets, goog_client_impl.EndpointGoogleTranslate, z.scopes, z.name, ctl)
	return
}

func (z *connTranslateImpl) PeerName() string {
	return z.name
}

func (z *connTranslateImpl) SetPeerName(name string) {
	z.name = name
}

func (z *connTranslateImpl) ScopeLabel() string {
	return app_definitions.ServiceGoogleTranslate
}

func (z *connTranslateImpl) ServiceName() string {
	return api_conn.ServiceGoogleTranslate
}

func (z *connTranslateImpl) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connTranslateImpl) Scopes() []string {
	return z.scopes
}

func (z *connTranslateImpl) Client() goog_client.Client {
	return z.ctx
}

func (z *connTranslateImpl) IsTranslate() bool {
	return true
}
