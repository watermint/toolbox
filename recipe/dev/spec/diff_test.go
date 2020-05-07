package spec

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDiff_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Diff{})
}
