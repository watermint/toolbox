package ig_bootstrap

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCleanup_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Autodelete{})
}
