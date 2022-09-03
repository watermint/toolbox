package api_auth_oauth

import (
	"errors"
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewSessionAlwaysFail(ctl app_control.Control) api_auth2.OAuthSession {
	return &sessionAlwaysFail{}
}

type sessionAlwaysFail struct {
}

func (z sessionAlwaysFail) Start(session api_auth2.OAuthSessionData) (entity api_auth2.OAuthEntity, err error) {
	return api_auth2.NewNoAuthOAuthEntity(), errors.New("always fail")
}
