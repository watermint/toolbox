package subtitles

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestOptimize_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Optimize{})
}
