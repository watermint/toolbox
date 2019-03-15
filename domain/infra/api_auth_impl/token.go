package api_auth_impl

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_context_impl"
	"github.com/watermint/toolbox/model/dbx_auth"
)

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

func Auth(ec *app.ExecContext, tokenType string) (ctx api_context.Context, err error) {
	au := dbx_auth.NewDefaultAuth(ec)
	legacyCtx, err := au.Auth(tokenType)
	if err != nil {
		return nil, err
	}
	return api_context_impl.New(ec, NewCompatible(legacyCtx.Token)), nil
}
