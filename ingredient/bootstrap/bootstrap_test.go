package bootstrap

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestBootstrap_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Bootstrap{})
}
