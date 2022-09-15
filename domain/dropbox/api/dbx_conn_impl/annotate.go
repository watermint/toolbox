package dbx_conn_impl

import (
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func newAnnotate(ctl app_control.Control, session api_auth.OAuthSession) api_auth.OAuthSession {
	return &annotateSession{
		ctl:     ctl,
		session: session,
	}
}

type annotateSession struct {
	ctl     app_control.Control
	session api_auth.OAuthSession
}

func (z annotateSession) annotateIndividual(entity api_auth.OAuthEntity) (annotated api_auth.OAuthEntity) {
	l := z.ctl.Log().With(esl.String("peerName", entity.PeerName), esl.Strings("scopes", entity.Scopes))
	client := dbx_client_impl.New(z.ctl, dbx_auth.DropboxIndividual, entity)
	res := client.Post("users/get_current_account")
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to verify token", esl.Error(err))
		return entity
	}

	j := res.Success().Json()
	displayName, found := j.FindString("name.display_name")
	if !found {
		l.Debug("Unable to find display name")
		return entity
	}
	email, found := j.FindString("email")
	if !found {
		entity.Description = fmt.Sprintf("%s", displayName)
		l.Debug("Unable to find email")
		return entity
	}
	l.Debug("Token Verified", esl.String("displayName", displayName))

	entity.Description = fmt.Sprintf("%s (%s)", displayName, email)
	return entity
}

func (z annotateSession) annotateTeam(entity api_auth.OAuthEntity) (annotated api_auth.OAuthEntity) {
	l := z.ctl.Log().With(esl.String("peerName", entity.PeerName), esl.Strings("scopes", entity.Scopes))
	client := dbx_client_impl.New(z.ctl, dbx_auth.DropboxTeam, entity)
	res := client.Post("team/get_info")
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to verify token", esl.Error(err))
		return entity
	}
	j := res.Success().Json()
	name, found := j.FindString("name")
	if !found {
		l.Debug("Team name not found")
		return entity
	}
	l.Debug("Token Verified", esl.String("name", name))
	entity.Description = name
	return entity
}

func (z annotateSession) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	entity, err = z.session.Start(session)
	if err != nil {
		return entity, err
	}

	if dbx_auth.HasTeamInfoRead(entity.Scopes) {
		return z.annotateTeam(entity), nil
	}
	if dbx_auth.HasAccountInfoRead(entity.Scopes) {
		return z.annotateIndividual(entity), nil
	}
	return entity, nil
}
