package desktop

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestSuspendupgrade_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Suspendupdate{})
}
