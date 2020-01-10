package batch

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestPermdelete_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Permdelete{})
}
