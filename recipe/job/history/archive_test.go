package history

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestArchive_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Archive{})
}