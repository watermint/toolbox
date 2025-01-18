package dbx_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func authSession(ctl app_control.Control) api_auth.OAuthSession {
	if ctl.Feature().Experiment(app_definitions.ExperimentDbxAuthRedirect) {
		return api_auth_oauth.NewSessionRedirect(ctl)
	} else if f, found := ctl.Feature().OptInGet(&api_auth_oauth.OptInFeatureRedirect{}); found && f.OptInIsEnabled() {
		return api_auth_oauth.NewSessionRedirect(ctl)
	} else {
		return api_auth_oauth.NewSessionCodeAuth(ctl)
	}
}

func connect(scopes []string, peerName string, ctl app_control.Control, app api_auth.OAuthAppData, isTeam bool) (ctx dbx_client.Client, err error) {
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
		s2 := newAnnotate(ctl, s1)
		s3 := api_auth_oauth.NewSessionRepository(s2, ctl.AuthRepository())

		entity, err := s3.Start(session)
		if err != nil {
			return nil, err
		}
		resolver := dbx_filesystem_impl.GetByEntity(dbx_client_impl.New(ctl, app, entity, dbx_filesystem_impl.NewEmpty()), entity)
		return dbx_client_impl.New(ctl, app, entity, resolver), nil
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
