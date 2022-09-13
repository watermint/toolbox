package as_conn

import (
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnAsanaApi interface {
	api_conn.ScopedConnection

	Client() as_client.Client
}
