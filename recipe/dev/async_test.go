package dev

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestAsync_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Async{})
}
