package paper

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAppend_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Append{})
}
