package translate

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestText_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Text{})
}
