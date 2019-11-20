package compare

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestAccount_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Account{})
}
