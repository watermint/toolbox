package file

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestMove_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Move{})
}
