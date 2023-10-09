package japanese

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestWakati_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Wakati{})
}
