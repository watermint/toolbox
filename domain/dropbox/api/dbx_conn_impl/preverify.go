package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/app"
)

type PreVerify interface {
	SetPreVerify(enabled bool)
}

func EnsurePreVerify(conn dbx_conn.ConnDropboxApi) {
	switch conn.PeerName() {
	case DefaultPeerName, app.PeerEndToEndTest:
		return

	default:
		if pv, ok := conn.(PreVerify); ok {
			pv.SetPreVerify(true)
		}
	}
}
