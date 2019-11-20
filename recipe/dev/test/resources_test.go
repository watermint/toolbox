package test

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestQuality_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Resources{})
}
