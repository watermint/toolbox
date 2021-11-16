package batch

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUnshare_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Unshare{})
}
