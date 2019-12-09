package diag

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestExplorer_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Explorer{})
}
