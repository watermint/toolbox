package auth

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestImport_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Import{})
}
