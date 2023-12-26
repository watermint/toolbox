package account

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFeature_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Feature{})
}
