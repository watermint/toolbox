package export

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFrame_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Frame{})
}
