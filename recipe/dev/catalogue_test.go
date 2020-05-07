package dev

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCatalogue_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Catalogue{})
}
