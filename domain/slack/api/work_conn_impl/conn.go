package work_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/slack/api/work_auth"
	"github.com/watermint/toolbox/domain/slack/api/work_client"
	"github.com/watermint/toolbox/domain/slack/api/work_client_impl"
	"github.com/watermint/toolbox/domain/slack/api/work_conn"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewSlackApi(name string) work_conn.ConnSlackApi {
	return &connSlackApi{
		name:   name,
		scopes: make([]string, 0),
	}
}

type connSlackApi struct {
	name   string
	ctx    work_client.Client
	scopes []string
}

func (z *connSlackApi) Connect(ctl app_control.Control) (err error) {
	ac, useMock, err := api_conn_impl.Connect(z.Scopes(), z.name, work_auth.New(ctl), ctl)
	if useMock {
		z.ctx = work_client_impl.NewMock(z.name, ctl)
		return nil
	}
	if ac != nil {
		z.ctx = work_client_impl.New(z.name, ctl, ac)
		return nil
	}
	if err != nil {
		return err
	} else {
		return errors.New("unknown state")
	}
}

func (z *connSlackApi) PeerName() string {
	return z.name
}

func (z *connSlackApi) SetPeerName(name string) {
	z.name = name
}

func (z *connSlackApi) ScopeLabel() string {
	return api_auth.Slack
}

func (z *connSlackApi) ServiceName() string {
	return api_conn.ServiceSlack
}

func (z *connSlackApi) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connSlackApi) Scopes() []string {
	return z.scopes
}

func (z *connSlackApi) Context() work_client.Client {
	return z.ctx
}
