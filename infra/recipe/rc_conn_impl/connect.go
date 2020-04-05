package rc_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
)

const (
	DefaultPeerName = "default"
)

func IsEndToEndTokenAllAvailable(ctl app_control.Control) bool {
	tts := []string{
		api_auth.DropboxTokenFull,
		api_auth.DropboxTokenBusinessInfo,
		api_auth.DropboxTokenBusinessFile,
		api_auth.DropboxTokenBusinessManagement,
		api_auth.DropboxTokenBusinessAudit,
	}

	for _, tt := range tts {
		_, err := ConnectTest(tt, qt_endtoend.EndToEndPeer, ctl)
		if err != nil {
			ctl.Log().Debug("Token is not available for the token type", zap.String("tokenType", tt), zap.Error(err))
			return false
		}
	}
	return true
}

func ConnectTest(tokenType, peerName string, ctl app_control.Control) (ctx api_context.DropboxApiContext, err error) {
	l := ctl.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	l.Debug("Connect for testing")
	if qt_endtoend.IsSkipEndToEndTest() {
		return nil, qt_errors.ErrorSkipEndToEndTest
	}
	peers := []string{peerName, qt_endtoend.EndToEndPeer}
	for _, peer := range peers {
		l.Debug("Retrieve cache from peer", zap.String("peer", peer))
		a := dbx_auth.NewConsoleCacheOnly(ctl, peer)
		if c, err := a.Auth(tokenType); err == nil {
			return dbx_context.New(ctl, c), nil
		}
	}
	return nil, qt_errors.ErrorNotEnoughResource
}

func connect(tokenType, peerName string, verify bool, ctl app_control.Control) (ctx api_context.DropboxApiContext, err error) {
	l := ctl.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	ui := ctl.UI()

	if c, ok := ctl.(app_control.ControlTestExtension); ok {
		if c.TestValue(qt_endtoend.CtlTestExtUseMock) == true {
			l.Debug("Test with mock")
			return dbx_context.NewMock(ctl), nil
		}
	}

	switch {
	case ctl.IsTest():
		if qt_endtoend.IsSkipEndToEndTest() {
			l.Debug("Skip end to end test")
			return dbx_context.NewMock(ctl), nil
		}
		return ConnectTest(tokenType, peerName, ctl)

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := dbx_auth_attr.NewConsole(ctl, peerName)
		ctx, err := c.Auth(tokenType)
		if err != nil {
			return nil, err
		}
		return dbx_context.New(ctl, ctx), nil

	case ui.IsWeb():
		l.Debug("Connect through web UI")
		a := dbx_auth_attr.NewWeb(ctl)
		tokens, err := a.List(tokenType)
		if err != nil {
			return nil, err
		}
		for _, t := range tokens {
			if t.PeerName() == peerName {
				c := dbx_context.New(ctl, t)
				return c, nil
			}
		}
		l.Debug("No peer found in existing connection")
		return nil, errors.New("no peer found")
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
