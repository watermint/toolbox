package file

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestMerge_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Merge{})
}
