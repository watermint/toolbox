package quota

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUpdate_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Update{})
}
