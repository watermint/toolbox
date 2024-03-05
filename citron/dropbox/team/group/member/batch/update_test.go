package bulk

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUpdate_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Update{})
}
