package hs_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_client"
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_client_impl"
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_basic"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/api/api_conn_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgDropboxSign struct {
	AskApiKey app_msg.Message
}

var (
	MHelloSign = app_msg.Apply(&MsgDropboxSign{}).(*MsgDropboxSign)
)

func NewConnHelloSign(name string) hs_conn.ConnHelloSignApi {
	return &connHelloSignApi{
		peerName: name,
	}
}

type connHelloSignApi struct {
	peerName string
	client   hs_client.Client
}

func (z *connHelloSignApi) IsBasic() bool {
	return true
}

func (z *connHelloSignApi) Connect(ctl app_control.Control) (err error) {
	l := ctl.Log()
	sessionData := api_auth.BasicSessionData{
		AppData: api_auth.BasicAppData{
			AppKeyName:      api_conn.ServiceDropboxSign,
			DontUseUsername: false,
			DontUsePassword: true,
		},
		PeerName: z.peerName,
	}
	entity, mock, err := api_conn_impl.BasicConnect(
		sessionData,
		ctl,
		api_auth_basic.CustomAskUserName(MHelloSign.AskApiKey),
	)
	if mock {
		z.client = hs_client_impl.NewMock(z.peerName, ctl)
		return nil
	}
	if err != nil {
		l.Debug("Unable to acquire", esl.Error(err))
		return err
	}
	z.client = hs_client_impl.New(z.peerName, ctl, entity)
	return nil
}

func (z *connHelloSignApi) PeerName() string {
	return z.peerName
}

func (z *connHelloSignApi) SetPeerName(name string) {
	z.peerName = name
}

func (z *connHelloSignApi) ScopeLabel() string {
	return app_definitions.ServiceDropboxSign
}

func (z *connHelloSignApi) ServiceName() string {
	return api_conn.ServiceDropboxSign
}

func (z *connHelloSignApi) Client() hs_client.Client {
	return z.client
}
