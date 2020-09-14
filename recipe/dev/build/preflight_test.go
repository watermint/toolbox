package build

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPreflight_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Preflight{})
}
