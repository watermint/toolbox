package tag

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCreate_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Create{})
}
