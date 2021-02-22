package sheet

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestClear_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Clear{})
}
