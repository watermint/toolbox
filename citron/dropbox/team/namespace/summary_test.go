package namespace

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestSummary_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Summary{})
}
