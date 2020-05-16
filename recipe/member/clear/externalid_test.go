package clear

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExternalid_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Externalid{})
}
