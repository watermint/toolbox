package uuid

import (
	"github.com/oklog/ulid/v2"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
)

type Ulid struct {
}

func (z *Ulid) Preset() {
}

func (z *Ulid) Exec(c app_control.Control) error {
	u := ulid.Make()
	ui_out.TextOut(c, u.String())
	return nil
}

func (z *Ulid) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Ulid{}, rc_recipe.NoCustomValues)
}
