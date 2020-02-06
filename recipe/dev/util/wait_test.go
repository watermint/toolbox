package util

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestWait_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Wait{})
}
