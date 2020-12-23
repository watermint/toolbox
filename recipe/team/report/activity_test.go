package report

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestActivity_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Activity{})
}
