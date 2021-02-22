package datetime

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestNow_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Now{})
}
