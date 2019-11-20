package dev

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestPreflight_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Preflight{})
}
