package setup

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMassfiles_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Massfiles{})
}
