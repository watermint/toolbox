package dev

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEcho_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Echo{})
}
