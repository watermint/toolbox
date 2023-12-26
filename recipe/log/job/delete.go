package job

import (
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/ig_job"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Delete struct {
	rc_recipe.RemarkConsole
	Days   mo_int.RangeInt
	Path   mo_string.OptionalString
	Delete *ig_job.Delete
}

func (z *Delete) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Delete, func(r rc_recipe.Recipe) {
		m := r.(*ig_job.Delete)
		m.Days = z.Days.Value()
		m.Path = z.Path
	})
}

func (z *Delete) Test(c app_control.Control) error {
	workspace, err := qt_file.MakeTestFolder("delete", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(workspace)
	}()

	return rc_exec.Exec(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Days.SetValue(365)
		m.Path = mo_string.NewOptional(workspace)
	})
}

func (z *Delete) Preset() {
	z.Days.SetRange(1, 3650, 28)
}
