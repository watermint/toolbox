package kvs

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDump_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Dump{})
}
