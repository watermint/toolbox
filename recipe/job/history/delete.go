package history

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/job"
)

type Delete struct {
	rc_recipe.RemarkConsole
	Days   mo_int.RangeInt
	Delete *job.Delete
}

func (z *Delete) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Delete, func(r rc_recipe.Recipe) {
		m := r.(*job.Delete)
		m.Days = z.Days.Value()
	})
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Days.SetValue(365)
	})
}

func (z *Delete) Preset() {
	z.Days.SetRange(1, 3650, 28)
}
