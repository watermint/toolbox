package activity

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestActivity_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Event{})
}
