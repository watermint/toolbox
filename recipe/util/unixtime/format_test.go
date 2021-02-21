package unixtime

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFormat_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Format{})
}
