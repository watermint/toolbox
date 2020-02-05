package export

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDoc_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Doc{})
}
