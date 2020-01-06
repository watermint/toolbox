package teamfolder

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestPermDelete_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Permdelete{})
}
