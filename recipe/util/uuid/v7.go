package uuid

import (
	"github.com/watermint/toolbox/essentials/strings/es_uuid"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"strings"
)

type V7 struct {
	UpperCase bool
}

func (z *V7) Preset() {
}

func (z *V7) Exec(c app_control.Control) error {
	v7 := es_uuid.NewV7()
	if z.UpperCase {
		ui_out.TextOut(c, strings.ToUpper(v7.String()))
	} else {
		ui_out.TextOut(c, v7.String())
	}
	return nil
}

func (z *V7) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &V7{}, rc_recipe.NoCustomValues)
}
