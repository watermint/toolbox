package goog_conn

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnGoogleApi interface {
	api_conn.ScopedConnection

	Context() goog_context.Context
}

type ConnGoogleMail interface {
	ConnGoogleApi
	IsGmail() bool
}
