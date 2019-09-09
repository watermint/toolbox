package app_run

import (
	"github.com/watermint/toolbox/infra/recpie/app_recipe_group"
	"github.com/watermint/toolbox/infra/recpie/app_test"
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
}
