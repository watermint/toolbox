package file

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCopy_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Copy{})
}
