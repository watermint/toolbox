package member

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestInvite_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Invite{})
}
