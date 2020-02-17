package connect

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestBusinessMgmt_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &BusinessMgmt{})
}
