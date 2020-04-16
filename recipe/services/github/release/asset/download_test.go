package asset

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDownload_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Download{})
}
