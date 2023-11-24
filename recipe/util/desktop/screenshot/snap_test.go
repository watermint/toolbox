package screenshot

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestSnap_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Snap{})
}
