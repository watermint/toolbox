package search

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestName_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Name{})
}
