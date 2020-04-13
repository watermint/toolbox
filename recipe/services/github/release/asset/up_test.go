package asset

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUp_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Up{})
}
