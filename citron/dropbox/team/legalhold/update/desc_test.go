package update

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDesc_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Desc{})
}
