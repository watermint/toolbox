package release

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDraft_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Draft{})
}
