package restore

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestRestore_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &All{})
}
