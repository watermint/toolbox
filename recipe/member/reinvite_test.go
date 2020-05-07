package member

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestReinvite_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Reinvite{})
}
