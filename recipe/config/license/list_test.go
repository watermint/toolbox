package license

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestList_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &List{})
}
