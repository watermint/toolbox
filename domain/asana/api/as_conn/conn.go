package as_conn

import (
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnAsanaApi interface {
	api_conn.ScopedConnection

	Context() as_client.Client
}
