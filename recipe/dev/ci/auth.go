package ci

import (
	"github.com/watermint/toolbox/infra/api/dbx_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Auth struct {
	Full  rc_conn.ConnUserFile
	Info  rc_conn.ConnBusinessInfo
	File  rc_conn.ConnBusinessFile
	Audit rc_conn.ConnBusinessAudit
	Mgmt  rc_conn.ConnBusinessMgmt
}

func (z *Auth) Preset() {
	z.Full.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Info.SetPeerName(qt_endtoend.EndToEndPeer)
	z.File.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Audit.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Mgmt.SetPeerName(qt_endtoend.EndToEndPeer)
}

func (z *Auth) Exec(c app_control.Control) error {
	if err := dbx_auth.CreateCompatible(c, qt_endtoend.EndToEndPeer); err != nil {
		return err
	}
	return nil
}

func (z *Auth) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
