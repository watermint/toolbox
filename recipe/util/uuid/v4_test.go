package uuid

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestV4_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &V4{})
}
