package insight

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestScanretry_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Scanretry{})
}