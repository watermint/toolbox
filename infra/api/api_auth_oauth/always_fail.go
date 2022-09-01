package api_auth_oauth

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewSessionAlwaysFail(ctl app_control.Control) api_auth.OAuthSession {
	return &sessionAlwaysFail{}
}

type sessionAlwaysFail struct {
}

func (z sessionAlwaysFail) Start(session api_auth.OAuthSessionData) (entity api_auth.OAuthEntity, err error) {
	return api_auth.NewNoAuthOAuthEntity(), errors.New("always fail")
}
