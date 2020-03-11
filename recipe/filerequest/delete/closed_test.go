package delete

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestClosed_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Closed{})
}
