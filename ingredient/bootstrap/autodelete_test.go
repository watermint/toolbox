package bootstrap

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCleanup_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Autodelete{})
}
