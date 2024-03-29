package teamfolder

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/ig_dropbox/ig_teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Replication struct {
	rc_recipe.RemarkIrreversible
	rc_recipe.RemarkExperimental
	Name        string
	Replication *ig_teamfolder.Replication
	SrcPeerName string
	DstPeerName string
}

func (z *Replication) Preset() {
	z.SrcPeerName = "src"
	z.DstPeerName = "dst"
}

func (z *Replication) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, &ig_teamfolder.Replication{}, func(r rc_recipe.Recipe) {
		rc := r.(*ig_teamfolder.Replication)
		rc.TargetNames = []string{z.Name}
		rc.Src.SetPeerName(z.SrcPeerName)
		rc.Dst.SetPeerName(z.DstPeerName)
	})
}

func (z *Replication) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		m := r.(*Replication)
		m.Name = "Sales"
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
