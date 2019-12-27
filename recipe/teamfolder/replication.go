package teamfolder

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Replication struct {
	Name        string
	Replication *teamfolder.Replication
}

func (z *Replication) Preset() {
}

func (z *Replication) Console() {
}

func (z *Replication) Exec(k rc_kitchen.Kitchen) error {
	return rc_exec.Exec(k.Control(), &teamfolder.Replication{}, func(r rc_recipe.Recipe) {
		rc := r.(*teamfolder.Replication)
		rc.TargetNames = []string{z.Name}
	})
}

func (z *Replication) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
}
