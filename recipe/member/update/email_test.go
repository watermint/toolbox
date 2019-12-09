package update

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEmail_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Email{})
}
