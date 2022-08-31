package dbx_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	DefaultPeerName = "default"
)

func authSession(ctl app_control.Control) api_auth.OAuthSession {
	if f, found := ctl.Feature().OptInGet(&api_auth_oauth.OptInFeatureRedirect{}); found && f.OptInIsEnabled() {
		return api_auth_oauth.NewSessionRedirect(ctl)
	} else {
		return api_auth_oauth.NewSessionCodeAuth(ctl)
	}
}

func connect(scopes []string, peerName string, ctl app_control.Control, app api_auth.OAuthAppData) (ctx dbx_client.Client, err error) {
	l := ctl.Log().With(esl.Strings("scopes", scopes), esl.String("peerName", peerName))
	ui := ctl.UI()

	session := api_auth.OAuthSessionData{
		AppData:  app,
		PeerName: peerName,
		Scopes:   scopes,
	}

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return dbx_client_impl.NewMock(peerName, ctl), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithReplay(); enabled {
		l.Debug("Test with replay")
		return dbx_client_impl.NewReplayMock(peerName, ctl, replay), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithSeqReplay(); enabled {
		l.Debug("Test with replay")
		return dbx_client_impl.NewSeqReplayMock(peerName, ctl, replay), nil
	}

	switch {
	case ctl.Feature().IsTest():
		l.Debug("Skip end to end test")
		return dbx_client_impl.NewMock(peerName, ctl), nil

	case ui.IsConsole():
		l.Debug("Connect through console UI")

		s1 := authSession(ctl)
		s2 := api_auth_oauth.NewSessionRepository(s1, ctl.AuthRepository())

		entity, err := s2.Start(session)
		if err != nil {
			return nil, err
		}
		return dbx_client_impl.New(ctl, app, entity), nil
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
