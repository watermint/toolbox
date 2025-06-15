package review

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestBatch_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Batch{})
}