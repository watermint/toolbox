package config

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDisable_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Disable{})
}
