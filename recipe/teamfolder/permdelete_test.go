package teamfolder

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPermDelete_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Permdelete{})
}
