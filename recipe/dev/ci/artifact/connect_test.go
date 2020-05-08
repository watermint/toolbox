package artifact

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestConnect_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Connect{})
}
