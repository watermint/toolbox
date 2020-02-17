package connect

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUserFile_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &UserFile{})
}
