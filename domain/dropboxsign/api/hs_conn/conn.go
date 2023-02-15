package hs_conn

import (
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnHelloSignApi interface {
	api_conn.BasicConnection

	Client() hs_client.Client
}
