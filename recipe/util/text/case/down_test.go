package _case

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDown_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Down{})
}
