package test

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestMonkey_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Monkey{})
}
