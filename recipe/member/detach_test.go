package member

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDetach_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Detach{})
}
