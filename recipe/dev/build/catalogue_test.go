package build

import (
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCatalogue_Exec(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}
	qtr_endtoend.TestRecipe(t, &Catalogue{})
}
