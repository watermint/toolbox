package cat

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestJobid_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Job{})
}
