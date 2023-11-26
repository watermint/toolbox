package desktop

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestOpen_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Open{})
}
