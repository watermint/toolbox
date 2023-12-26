package feature

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFeatures_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &List{})
}
