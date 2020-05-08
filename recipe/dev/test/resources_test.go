package test

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestQuality_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Resources{})
}
