package review

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestOptions_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Options{})
}
