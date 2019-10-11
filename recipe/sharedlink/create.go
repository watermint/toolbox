package sharedlink

import "github.com/watermint/toolbox/infra/recpie/app_conn"

type CreateVO struct {
	Peer app_conn.ConnUserFile
	Path string
}
