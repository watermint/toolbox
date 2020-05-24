package content

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestGet_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Get{})
}
