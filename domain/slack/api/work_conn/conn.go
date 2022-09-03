package work_conn

import (
	"github.com/watermint/toolbox/domain/slack/api/work_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnSlackApi interface {
	api_conn.ScopedConnection

	Context() work_client.Client
}
