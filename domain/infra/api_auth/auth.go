package api_auth

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
)

type Auth interface {
	Auth(tokenType string) (ctx api_context.Context, err error)
}

type Token interface {
	Token() string
}
