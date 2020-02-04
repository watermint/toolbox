package member

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestAdd_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Add{})
}
