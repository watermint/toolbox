package template

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestApply_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Apply{})
}
