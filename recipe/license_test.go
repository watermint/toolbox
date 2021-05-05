package recipe

import (
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestLicense_Exec(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}
	qtr_endtoend.TestRecipe(t, &License{})
}
