package util

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAnon_Anonymise(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Anonymise{})
}
