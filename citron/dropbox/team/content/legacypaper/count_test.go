package legacypaper

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCount_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Count{})
}
