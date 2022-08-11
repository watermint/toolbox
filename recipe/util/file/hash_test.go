package file

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestHash_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Hash{})
}
