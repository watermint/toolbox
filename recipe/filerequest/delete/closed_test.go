package delete

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestClosed_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Closed{})
}
