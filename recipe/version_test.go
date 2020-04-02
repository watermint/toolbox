package recipe

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestVersion_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Version{})
}
