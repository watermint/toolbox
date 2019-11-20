package test

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRecipe_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Recipe{})
}
