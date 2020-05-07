package update

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExternalId_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Externalid{})
}
