package log

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestKind_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Kind{})
}
