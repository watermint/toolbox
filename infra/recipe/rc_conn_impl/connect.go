package rc_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
)

const (
	DefaultPeerName = "default"
)

func ConnectTest(tokenType, peerName string, ctl app_control.Control) (ctx api_context.Context, err error) {
	l := ctl.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	l.Debug("Connect for testing")
	if qt_endtoend.IsSkipEndToEndTest() {
		return nil, qt_errors.ErrorSkipEndToEndTest
	}
	a := api_auth_impl.NewCached(ctl, api_auth_impl.PeerName(peerName))
	if c, err := a.Auth(tokenType); err == nil {
		return c, nil
	}

	// fallback to end to end peer
	a = api_auth_impl.NewCached(ctl, api_auth_impl.PeerName(qt_endtoend.EndToEndPeer))
	if c, err := a.Auth(tokenType); err == nil {
		return c, nil
	} else {
		return nil, qt_errors.ErrorNotEnoughResource
	}
}

func connect(tokenType, peerName string, verify bool, ctl app_control.Control) (ctx api_context.Context, err error) {
	l := ctl.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	ui := ctl.UI()

	verifyToken := func(ctx0 api_context.Context) error {
		if verify {
			desc, suppl, err0 := api_auth_impl.VerifyToken(tokenType, ctx0)
			if err0 == nil {
				ui.Info(MConnect.VerifySuccess.With("Desc", desc).With("Suppl", suppl))
			}
			return err0
		}
		return nil
	}

	switch {
	case ctl.IsTest():
		if c, ok := ctl.(app_control.ControlTestExtension); ok {
			if c.TestValue(qt_endtoend.CtlTestExtUseMock) == true {
				l.Debug("Test with mock")
				return api_context_impl.NewMock(ctl), nil
			}
		}
		if qt_endtoend.IsSkipEndToEndTest() {
			l.Debug("Skip end to end test")
			return api_context_impl.NewMock(ctl), nil
		}
		return ConnectTest(tokenType, peerName, ctl)

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := api_auth_impl.New(ctl, api_auth_impl.PeerName(peerName))
		ctx, err = c.Auth(tokenType)
		if err == nil {
			err = verifyToken(ctx)
		}
		return

	case ui.IsWeb():
		l.Debug("Connect through web UI")
		a := api_auth_impl.NewWeb(ctl)
		tokens, err := a.List(tokenType)
		if err != nil {
			return nil, err
		}
		for _, t := range tokens {
			if t.PeerName == peerName {
				c := api_context_impl.New(ctl, t)
				return c, nil
			}
		}
		l.Debug("No peer found in existing connection")
		return nil, errors.New("no peer found")
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
