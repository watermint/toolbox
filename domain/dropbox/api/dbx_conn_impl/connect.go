package dbx_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
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
		helper, err := getHelper(ctl, entity, isTeam)
		if err != nil {
			return nil, err
		}
		return dbx_client_impl.New(ctl, app, entity, helper), nil
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}

func getHelper(ctl app_control.Control, entity api_auth.OAuthEntity, isTeam bool) (helper *dbx_filesystem.FileSystemBuilderHelper, err error) {
	if isTeam {
		return getHelperTeam(ctl, entity)
	} else {
		return getHelperIndividual(ctl, entity)
	}
}

func getHelperIndividual(ctl app_control.Control, entity api_auth.OAuthEntity) (helper *dbx_filesystem.FileSystemBuilderHelper, err error) {
	l := ctl.Log().With(esl.String("peerName", entity.PeerName), esl.Strings("scopes", entity.Scopes))
	client := dbx_client_impl.New(ctl, dbx_auth.DropboxIndividual, entity, dbx_filesystem.NewEmptyHelper())

	res := client.Post("users/get_current_account")
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to verify token", esl.Error(err))
		return nil, err
	}
	rj := res.Success().Json()
	rootInfo := &dbx_filesystem.RootInfo{}
	if rj.Model(rootInfo) != nil {
		l.Debug("Unable to find root info")
		return nil, errors.New("unable to find root info")
	}

	return dbx_filesystem.NewHelper(rootInfo), nil
}

func getHelperTeam(ctl app_control.Control, entity api_auth.OAuthEntity) (helper *dbx_filesystem.FileSystemBuilderHelper, err error) {
	l := ctl.Log().With(esl.String("peerName", entity.PeerName), esl.Strings("scopes", entity.Scopes))
	client := dbx_client_impl.New(ctl, dbx_auth.DropboxIndividual, entity, dbx_filesystem.NewEmptyHelper())

	res := client.Post("team/token/get_authenticated_admin")
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to verify token", esl.Error(err))
		return nil, err
	}
	rj := res.Success().Json()
	memberNamespaceId, ok := rj.FindString("admin_profile.member_folder_id")
	if !ok {
		l.Debug("Unable to find team folder id")
		return nil, errors.New("unable to find team folder id")
	}
	rootNamespaceId, ok := rj.FindString("admin_profile.root_folder_id")
	if !ok {
		l.Debug("Unable to find root namespace id")
		return nil, errors.New("unable to find root namespace id")
	}
	return dbx_filesystem.NewHelper(&dbx_filesystem.RootInfo{
		HomeNamespaceId: memberNamespaceId,
		RootNamespaceId: rootNamespaceId,
	}), nil
}
