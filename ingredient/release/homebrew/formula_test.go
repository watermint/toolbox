package homebrew

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFormula_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Formula{})
}
