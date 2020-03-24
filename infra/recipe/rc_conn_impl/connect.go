package rc_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/dbx_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
)

const (
	DefaultPeerName = "default"
)

var (
	ErrorIncompatibleTokenType = errors.New("incompatible token type")
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
	a := dbx_auth.NewCached(ctl, dbx_auth.PeerName(peerName))
	if c, err := a.Auth(tokenType); err == nil {
		if dc, ok := c.(api_context.DropboxApiContext); ok {
			return dc, nil
		} else {
			l.Debug("Incompatible token type found", zap.Any("token", c))
			return nil, ErrorIncompatibleTokenType
		}
	}

	// fallback to end to end peer
	a = dbx_auth.NewCached(ctl, dbx_auth.PeerName(qt_endtoend.EndToEndPeer))
	if c, err := a.Auth(tokenType); err == nil {
		if dc, ok := c.(api_context.DropboxApiContext); ok {
			return dc, nil
		} else {
			l.Debug("Incompatible token type found", zap.Any("returned token", c))
			return nil, ErrorIncompatibleTokenType
		}
	} else {
		return nil, qt_errors.ErrorNotEnoughResource
	}
}

func connect(tokenType, peerName string, verify bool, ctl app_control.Control) (ctx api_context.DropboxApiContext, err error) {
	l := ctl.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	ui := ctl.UI()

	verifyToken := func(ctx0 api_context.DropboxApiContext) error {
		if verify {
			desc, suppl, err0 := dbx_auth.VerifyToken(tokenType, ctx0)
			if err0 == nil {
				ui.Info(MConnect.VerifySuccess.With("Desc", desc).With("Suppl", suppl))
			}
			return err0
		}
		return nil
	}

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
		c := dbx_auth.New(ctl, dbx_auth.PeerName(peerName))
		ctx0, err := c.Auth(tokenType)
		if err != nil {
			return nil, err
		}
		ctx, ok := ctx0.(api_context.DropboxApiContext)
		if !ok {
			l.Debug("Incompatible token type found", zap.Any("token", ctx0))
			return nil, ErrorIncompatibleTokenType
		}
		err = verifyToken(ctx)
		return ctx, nil

	case ui.IsWeb():
		l.Debug("Connect through web UI")
		a := dbx_auth.NewWeb(ctl)
		tokens, err := a.List(tokenType)
		if err != nil {
			return nil, err
		}
		for _, t := range tokens {
			if t.PeerName == peerName {
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
