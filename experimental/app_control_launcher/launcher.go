package app_control_launcher

import (
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_workspace"
)

type ControlLauncher interface {
	Catalogue() []app_recipe.Recipe
	NewControl(user app_workspace.MultiUser) app_control.Control
}
