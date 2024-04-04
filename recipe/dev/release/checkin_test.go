package release

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCheckin_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Checkin{})
}
