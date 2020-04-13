package auth

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestExport_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Export{})
}
