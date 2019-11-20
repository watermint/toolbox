package filerequest

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestClone_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Clone{})
}
