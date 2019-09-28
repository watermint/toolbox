package app_run

import (
	"flag"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_recipe_group"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"testing"
)

func TestCatalogue(t *testing.T) {
	_, _, _, ui := app_test.TestResources(t)
	testGroup(Catalogue(), ui)
}

func testGroup(g *app_recipe_group.Group, ui app_ui.UI) {
	g.PrintUsage(ui)
	for _, sg := range g.SubGroups {
		testGroup(sg, ui)
	}
	for _, r := range g.Recipes {
		testRecipe(g, r, ui)
	}
}

func testRecipe(g *app_recipe_group.Group, r app_recipe.Recipe, ui app_ui.UI) {
	vo := r.Requirement()
	f := flag.NewFlagSet("", flag.ContinueOnError)
	vc := app_vo_impl.NewValueContainer(vo)
	vc.MakeFlagSet(f, ui)
	g.PrintRecipeUsage(ui, r, f)
}
