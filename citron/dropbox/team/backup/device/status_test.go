package device

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestStatus_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Status{})
}
