package smoothie

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestVersion_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Flavor{})
}
