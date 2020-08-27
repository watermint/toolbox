package test

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAsync_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Async{})
}
