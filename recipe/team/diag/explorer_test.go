package diag

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExplorer_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Explorer{})
}
