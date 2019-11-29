package activity

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUser_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &User{})
}
