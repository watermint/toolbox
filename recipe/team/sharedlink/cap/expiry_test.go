package cap

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExpiry_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Expiry{})
}
