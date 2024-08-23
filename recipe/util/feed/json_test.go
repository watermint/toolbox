package feed

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUrl_Exec(t *testing.T) {
	if testing.Short() {
		return
	}
	qtr_endtoend.TestRecipe(t, &Json{})
}
