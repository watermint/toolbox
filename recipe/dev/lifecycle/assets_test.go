package lifecycle

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAssets_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Assets{})
}
