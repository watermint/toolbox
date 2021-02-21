package date

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestToday_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Today{})
}
