package quota

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUsage_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Usage{})
}
