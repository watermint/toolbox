package dbx_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/essentials/log/es_log"
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
		_, err := ConnectTest(tt, app.PeerEndToEndTest, ctl)
		if err != nil {
			ctl.Log().Debug("Token is not available for the token type", es_log.String("tokenType", tt), es_log.Error(err))
			return false
		}
	}
	return true
}

func ConnectTest(tokenType, peerName string, ctl app_control.Control) (ctx dbx_context.Context, err error) {
	l := ctl.Log().With(es_log.String("tokenType", tokenType), es_log.String("peerName", peerName))
	l.Debug("Connect for testing")
	if qt_endtoend.IsSkipEndToEndTest() {
		return nil, qt_errors.ErrorSkipEndToEndTest
	}
	peers := []string{peerName, app.PeerEndToEndTest}
	for _, peer := range peers {
		l.Debug("Retrieve cache from peer", es_log.String("peer", peer))
		a := api_auth_impl.NewConsoleCacheOnly(ctl, peer)
		if c, err := a.Auth(tokenType); err == nil {
			return dbx_context_impl.New(ctl, c), nil
		}
	}
	return nil, qt_errors.ErrorNotEnoughResource
}

func connect(tokenType, peerName string, verify bool, ctl app_control.Control) (ctx dbx_context.Context, err error) {
	l := ctl.Log().With(es_log.String("tokenType", tokenType), es_log.String("peerName", peerName))
	ui := ctl.UI()

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return dbx_context_impl.NewMock(ctl), nil
	}

	switch {
	case ctl.Feature().IsTest():
		if qt_endtoend.IsSkipEndToEndTest() {
			l.Debug("Skip end to end test")
			return dbx_context_impl.NewMock(ctl), nil
		}
		return ConnectTest(tokenType, peerName, ctl)

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := dbx_auth_attr.NewConsole(ctl, peerName)
		ctx, err := c.Auth(tokenType)
		if err != nil {
			return nil, err
		}
		return dbx_context_impl.New(ctl, ctx), nil

	case ui.IsWeb():
		l.Debug("Connect through web UI")
		a := dbx_auth_attr.NewWeb(ctl)
		tokens, err := a.List(tokenType)
		if err != nil {
			return nil, err
		}
		for _, t := range tokens {
			if t.PeerName() == peerName {
				c := dbx_context_impl.New(ctl, t)
				return c, nil
			}
		}
		l.Debug("No peer found in existing connection")
		return nil, errors.New("no peer found")
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
