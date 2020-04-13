package config

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestFeatures_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Features{})
}
