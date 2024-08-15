package deepl_conn_impl

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_client"
	"github.com/watermint/toolbox/domain/deepl/api/deepl_client_impl"
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/api/api_conn_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgDeeplApi struct {
	AskApiKey app_msg.Message
}

var (
	MDeeplApi = app_msg.Apply(&MsgDeeplApi{}).(*MsgDeeplApi)
)

func NewConnDeeplApi(name string) deepl_conn.ConnDeeplApi {
	return &connDeeplApiImpl{
		peerName: name,
	}
}

type connDeeplApiImpl struct {
	peerName string
	client   deepl_client.Client
}

func (z *connDeeplApiImpl) Connect(ctl app_control.Control) (err error) {
	l := ctl.Log()
	sessionData := api_auth.KeySessionData{
		AppData: api_auth.KeyAppData{
			AppKeyName: api_conn.ScopeLabelDeepl,
		},
		PeerName: z.peerName,
	}
	entity, mock, err := api_conn_impl.KeyConnect(
		sessionData,
		ctl,
		MDeeplApi.AskApiKey,
	)
	if mock {
		z.client = deepl_client_impl.NewMock(z.peerName, ctl)
		return nil
	}
	if err != nil {
		l.Debug("Unable to acquire", esl.Error(err))
		return err
	}
	z.client = deepl_client_impl.NewV2(z.peerName, ctl, entity)
	return nil
}

func (z *connDeeplApiImpl) PeerName() string {
	return z.peerName
}

func (z *connDeeplApiImpl) SetPeerName(name string) {
	z.peerName = name
}

func (z *connDeeplApiImpl) ScopeLabel() string {
	return app_definitions.ServiceKeyDeepl
}

func (z *connDeeplApiImpl) AppKeyName() string {
	return api_conn.ScopeLabelDeepl
}

func (z *connDeeplApiImpl) IsKeyAuth() bool {
	return true
}

func (z *connDeeplApiImpl) Client() deepl_client.Client {
	return z.client
}
