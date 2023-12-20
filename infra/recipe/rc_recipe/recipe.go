package rc_recipe

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

const (
	BasePackage = app_definitions.Pkg + "/recipe"
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
