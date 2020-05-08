package content

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPolicy_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Policy{})
}
