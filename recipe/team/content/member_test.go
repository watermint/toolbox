package content

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestMember_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Member{})
}
