package history

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestShip_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Ship{})
}
