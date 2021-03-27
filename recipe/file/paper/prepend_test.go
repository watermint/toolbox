package paper

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPrepend_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Prepend{})
}
