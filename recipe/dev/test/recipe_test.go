package test

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestRecipe_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Recipe{})
}
