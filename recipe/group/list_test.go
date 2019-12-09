package group

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestList_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &List{})
}
