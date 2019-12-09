package device

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUnlink_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Unlink{})
}
