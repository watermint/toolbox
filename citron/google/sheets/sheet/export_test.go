package sheet

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExport_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Export{})
}
