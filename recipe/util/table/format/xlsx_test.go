package format

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestXlsx_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Xlsx{})
}
