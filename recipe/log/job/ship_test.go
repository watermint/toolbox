package job

import (
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestShip_Exec(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}
	qtr_endtoend.TestRecipe(t, &Ship{})
}
