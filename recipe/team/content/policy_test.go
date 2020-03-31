package content

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestPolicy_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Policy{})
}
