package batch

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUrl_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Url{})
}
