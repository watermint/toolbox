package gui

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestGui_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Launch{})
}
