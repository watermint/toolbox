package spec

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDiff_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Diff{})
}
