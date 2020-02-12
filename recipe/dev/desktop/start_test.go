package desktop

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestLaunch_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Start{})
}
