package _import

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestViaApp_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &ViaApp{})
}
