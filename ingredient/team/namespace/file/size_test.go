package file

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestSize_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Size{})
}
