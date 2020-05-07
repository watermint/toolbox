package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	infra_recipe_rc_recipe "github.com/watermint/toolbox/infra/recipe/rc_recipe"
	ingredientbootstrap "github.com/watermint/toolbox/ingredient/bootstrap"
	ingredientfile "github.com/watermint/toolbox/ingredient/file"
	ingredientjob "github.com/watermint/toolbox/ingredient/job"
	ingredientteamnamespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	ingredientteamfolder "github.com/watermint/toolbox/ingredient/teamfolder"
)

func AutoDetectedIngredients() []infra_recipe_rc_recipe.Recipe {
	return []infra_recipe_rc_recipe.Recipe{
		&ingredientbootstrap.Autodelete{},
		&ingredientbootstrap.Bootstrap{},
		&ingredientfile.Upload{},
		&ingredientjob.Delete{},
		&ingredientteamnamespacefile.List{},
		&ingredientteamnamespacefile.Size{},
		&ingredientteamfolder.Replication{},
	}
}
