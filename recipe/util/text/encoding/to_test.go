package encoding

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestTo_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &To{})
}
