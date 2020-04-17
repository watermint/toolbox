package auth

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Connect struct {
	rc_recipe.RemarkSecret
	Full   dbx_conn.ConnUserFile
	Info   dbx_conn.ConnBusinessInfo
	File   dbx_conn.ConnBusinessFile
	Audit  dbx_conn.ConnBusinessAudit
	Mgmt   dbx_conn.ConnBusinessMgmt
	Github gh_conn.ConnGithubRepo
}

func (z *Connect) Preset() {
	z.Full.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Info.SetPeerName(qt_endtoend.EndToEndPeer)
	z.File.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Audit.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Mgmt.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Github.SetPeerName(qt_endtoend.DeployPeer)
}

func (z *Connect) Exec(c app_control.Control) error {
	return nil
}

func (z *Connect) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Connect{}, rc_recipe.NoCustomValues)
}
