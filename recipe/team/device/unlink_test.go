package device

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUnlink_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Unlink{})
}
