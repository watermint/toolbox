package database

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExec_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Exec{})
}
