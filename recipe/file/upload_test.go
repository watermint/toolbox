package file

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUpload_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Upload{})
}
