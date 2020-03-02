package member

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestReinvite_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Reinvite{})
}
