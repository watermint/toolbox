package job

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestHistory_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &History{})
}
