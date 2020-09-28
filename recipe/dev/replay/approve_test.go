package replay

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestApprove_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Approve{})
}
