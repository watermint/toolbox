package spec

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestSpec_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Doc{})
}
