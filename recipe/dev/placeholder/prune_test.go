package placeholder

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPrune_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Prune{})
}
