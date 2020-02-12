package desktop

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestStop_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Stop{})
}
