package file

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDelete_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Delete{})
}
