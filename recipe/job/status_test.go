package job

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestStatus_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Status{})
}
