package teamfolder

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Replication struct {
	Name        string
	Replication *teamfolder.Replication
	SrcPeerName string
	DstPeerName string
}

func (z *Replication) Preset() {
	z.SrcPeerName = "src"
	z.DstPeerName = "dst"
}

func (z *Replication) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, &teamfolder.Replication{}, func(r rc_recipe.Recipe) {
		rc := r.(*teamfolder.Replication)
		rc.TargetNames = []string{z.Name}
		rc.SrcFile.SetPeerName(z.SrcPeerName)
		rc.SrcMgmt.SetPeerName(z.SrcPeerName)
		rc.DstFile.SetPeerName(z.DstPeerName)
		rc.DstMgmt.SetPeerName(z.DstPeerName)
	})
}

func (z *Replication) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
