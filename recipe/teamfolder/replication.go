package teamfolder

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
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
	err := rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		m := r.(*Replication)
		m.Name = "Sales"
	})
	if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
