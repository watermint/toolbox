package update

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestExternalId_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Externalid{})
}
