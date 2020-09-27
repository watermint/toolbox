package dbx_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
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
		_, err := ConnectTest([]string{tt}, app.PeerEndToEndTest, ctl)
		if err != nil {
			ctl.Log().Debug("Token is not available for the token type", esl.String("tokenType", tt), esl.Error(err))
			return false
		}
	}
	return true
}

func ConnectTest(scopes []string, peerName string, ctl app_control.Control) (ctx dbx_context.Context, err error) {
	l := ctl.Log().With(esl.Strings("scopes", scopes), esl.String("peerName", peerName))
	l.Debug("Connect for testing")
	if qt_endtoend.IsSkipEndToEndTest() {
		return nil, qt_errors.ErrorSkipEndToEndTest
	}
	peers := []string{peerName, app.PeerEndToEndTest}
	for _, peer := range peers {
		l.Debug("Retrieve cache from peer", esl.String("peer", peer))
		a := api_auth_impl.NewConsoleCacheOnly(ctl, peer, dbx_auth.NewLegacyApp(ctl))
		if c, err := a.Auth(scopes); err == nil {
			return dbx_context_impl.New(ctl, c), nil
		}
	}
	return nil, qt_errors.ErrorNotEnoughResource
}

func connect(scopes []string, peerName string, ctl app_control.Control, app api_auth.App) (ctx dbx_context.Context, err error) {
	l := ctl.Log().With(esl.Strings("scopes", scopes), esl.String("peerName", peerName))
	ui := ctl.UI()

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return dbx_context_impl.NewMock(ctl), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithReplay(); enabled {
		l.Debug("Test with replay")
		return dbx_context_impl.NewReplayMock(ctl, replay), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithSeqReplay(); enabled {
		l.Debug("Test with replay")
		return dbx_context_impl.NewSeqReplayMock(ctl, replay), nil
	}

	switch {
	case ctl.Feature().IsTest():
		if qt_endtoend.IsSkipEndToEndTest() {
			l.Debug("Skip end to end test")
			return dbx_context_impl.NewMock(ctl), nil
		}
		return ConnectTest(scopes, peerName, ctl)

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := dbx_auth_attr.NewConsole(ctl, peerName, app)
		ctx, err := c.Auth(scopes)
		if err != nil {
			return nil, err
		}
		return dbx_context_impl.New(ctl, ctx), nil

	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
