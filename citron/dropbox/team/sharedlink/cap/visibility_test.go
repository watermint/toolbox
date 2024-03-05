package cap

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestVisibility_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Visibility{})
}
