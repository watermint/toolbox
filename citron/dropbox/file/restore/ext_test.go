package restore

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExt_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Ext{})
}
