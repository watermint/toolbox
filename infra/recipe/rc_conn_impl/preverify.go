package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type PreVerify interface {
	SetPreVerify(enabled bool)
}

func EnsurePreVerify(conn rc_conn.ConnDropboxApi) {
	switch conn.Name() {
	case DefaultPeerName, qt_endtoend.EndToEndPeer:
		return

	default:
		if pv, ok := conn.(PreVerify); ok {
			pv.SetPreVerify(true)
		}
	}
}
