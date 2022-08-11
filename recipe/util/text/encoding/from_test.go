package encoding

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFrom_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &From{})
}
