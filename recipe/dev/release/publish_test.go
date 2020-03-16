package release

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestPublish_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Publish{})
}
