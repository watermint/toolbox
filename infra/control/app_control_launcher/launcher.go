package app_control_launcher

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
)

type ControlLauncher interface {
	Catalogue() []app_recipe.Recipe
	NewControl(user app_workspace.MultiUser) (ctl app_control.Control, err error)
}
