package search

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestContent_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Content{})
}
