package rc_recipe

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

type Preset interface {
	Preset()
}

type Recipe interface {
	Preset
	Exec(c app_control.Control) error
	Test(c app_control.Control) error
}

func NoCustomValues(r Recipe) {}
