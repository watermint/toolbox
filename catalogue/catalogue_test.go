package catalogue

import (
	"flag"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCatalogue(t *testing.T) {
	_, _, _, ui := qt_recipe.Resources(t)
	testGroup(Groups(), ui)
	for _, r := range Ingredients() {
		spec := rc_spec.New(r)
		if spec == nil {
			continue
		}
		for _, m := range spec.Messages() {
			ui.Info(m)
		}
	}
}

func testGroup(g *rc_group.Group, ui app_ui.UI) {
	g.PrintUsage(ui)
	for _, sg := range g.SubGroups {
		testGroup(sg, ui)
	}
	for _, r := range g.Recipes {
		testRecipe(g, r, ui)
	}
}

func testRecipe(g *rc_group.Group, r rc_recipe.Recipe, ui app_ui.UI) {
	f := flag.NewFlagSet("", flag.ContinueOnError)
	spec := rc_spec.New(r)
	if spec == nil {
		// skip
		return
	}
	spec.SetFlags(f, ui)
	g.PrintRecipeUsage(ui, spec, f)

	for _, m := range spec.Messages() {
		ui.Info(m)
	}
}
