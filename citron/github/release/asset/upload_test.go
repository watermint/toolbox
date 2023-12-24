package asset

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUp_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Upload{})
}
