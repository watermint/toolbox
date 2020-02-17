package connect

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestBusinessInfo_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &BusinessInfo{})
}
