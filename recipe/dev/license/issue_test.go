package license

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestIssue_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Issue{})
}
