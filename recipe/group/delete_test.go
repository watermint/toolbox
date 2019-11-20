package group

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRemove_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Delete{})
}
