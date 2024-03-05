package member

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUnsuspend_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Unsuspend{})
}
