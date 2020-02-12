package search

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestName_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Name{})
}
