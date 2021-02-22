package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/api/goog_context_impl"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func connect(endpointBase goog_context_impl.EndpointType, scopes []string, name string, ctl app_control.Control) (ctx goog_context.Context, err error) {
	ac, useMock, err := api_conn_impl.Connect(scopes, name, goog_auth.NewApp(ctl), ctl)
	if useMock {
		return goog_context_impl.NewMock(endpointBase, name, ctl), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithSeqReplay(); enabled {
		return goog_context_impl.NewReplayMock(endpointBase, name, ctl, replay), nil
	}
	if ac != nil {
		return goog_context_impl.New(endpointBase, name, ctl, ac), nil
	}
	return nil, err
}
