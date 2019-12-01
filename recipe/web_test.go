package recipe

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestWeb_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Web{})
}
