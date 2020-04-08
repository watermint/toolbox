package config

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEnable_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Enable{})
}
