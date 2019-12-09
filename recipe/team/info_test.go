package team

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestInfo_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Info{})
}
