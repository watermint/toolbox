package group

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRename_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Rename{})
}
