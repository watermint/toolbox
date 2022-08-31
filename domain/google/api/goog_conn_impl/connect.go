package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/api/goog_client_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func connect(app api_auth.OAuthAppData, endpointBase goog_client_impl.EndpointType, scopes []string, peerName string, ctl app_control.Control) (ctx goog_client.Client, err error) {
	session := api_auth.OAuthSessionData{
		AppData:  app,
		PeerName: peerName,
		Scopes:   scopes,
	}
	entity, useMock, err := api_conn_impl.ConnectByRedirect(session, ctl)
	if useMock {
		return goog_client_impl.NewMock(endpointBase, peerName, ctl), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithSeqReplay(); enabled {
		return goog_client_impl.NewReplayMock(endpointBase, peerName, ctl, replay), nil
	}
	if err != nil {
		return nil, err
	}
	return goog_client_impl.New(endpointBase, peerName, ctl, entity), nil
}
