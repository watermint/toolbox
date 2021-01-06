package report

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDevices_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Devices{})
}
