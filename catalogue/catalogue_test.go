package catalogue

import (
	"flag"
	"github.com/watermint/toolbox/infra/recpie/rc_group"
	"github.com/watermint/toolbox/infra/recpie/rc_recipe"
	"github.com/watermint/toolbox/infra/recpie/rc_vo_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCatalogue(t *testing.T) {
	_, _, _, ui := qt_recipe.Resources(t)
	testGroup(Catalogue(), ui)
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
	switch scr := r.(type) {
	case rc_recipe.SideCarRecipe:
		vo := scr.Requirement()
		f := flag.NewFlagSet("", flag.ContinueOnError)
		vc := rc_vo_impl.NewValueContainer(vo)
		vc.MakeFlagSet(f, ui)
		g.PrintRecipeUsage(ui, r, f)
	}
}
