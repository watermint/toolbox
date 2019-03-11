package api_auth_impl

import "github.com/watermint/toolbox/domain/infra/api_auth"

type Compatible struct {
	token string
}

func (z *Compatible) Token() string {
	return z.token
}

func NewCompatible(token string) api_auth.Token {
	return &Compatible{
		token: token,
	}
}
