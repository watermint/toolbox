package uuid

import (
	"github.com/watermint/essentials/eformat/euuid"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"strings"
)

type V4 struct {
	UpperCase bool
}

func (z *V4) Preset() {
}

func (z *V4) Exec(c app_control.Control) error {
	v4 := euuid.NewV4()
	if z.UpperCase {
		ui_out.TextOut(c, strings.ToUpper(v4.String()))
	} else {
		ui_out.TextOut(c, v4.String())
	}
	return nil
}

func (z *V4) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &V4{}, rc_recipe.NoCustomValues)
}
