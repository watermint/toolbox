package api

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestJob_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Job{})
}
