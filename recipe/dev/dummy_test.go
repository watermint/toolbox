package dev

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDummy_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Dummy{})
}
