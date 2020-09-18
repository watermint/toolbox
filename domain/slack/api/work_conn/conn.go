package work_conn

import (
	"github.com/watermint/toolbox/domain/slack/api/work_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnSlackApi interface {
	api_conn.ScopedConnection

	Context() work_context.Context
}
