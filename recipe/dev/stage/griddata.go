package stage

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Griddata struct {
	rc_recipe.RemarkSecret
	In  da_griddata.GridDataInput
	Out da_griddata.GridDataOutput
}

func (z *Griddata) Preset() {
}

func (z *Griddata) Exec(c app_control.Control) error {
	return z.In.EachRow(func(col []interface{}, rowIndex int) error {
		z.Out.Row(col)
		return nil
	})
}

func (z *Griddata) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("grid_data", "alex@example.com,Alex\ndavid@example.com,David\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Griddata{}, func(r rc_recipe.Recipe) {
		m := r.(*Griddata)
		m.In.SetFilePath(f)
	})
}
