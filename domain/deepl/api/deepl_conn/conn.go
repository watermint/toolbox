package deepl_conn

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnDeeplApi interface {
	api_conn.KeyConnection

	Client() deepl_client.Client
}
