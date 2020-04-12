package artifact

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestConnect_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Connect{})
}
