package fg_conn

import (
	"github.com/watermint/toolbox/domain/figma/api/fg_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnFigmaApi interface {
	api_conn.BasicConnection

	Client() fg_client.Client
}
