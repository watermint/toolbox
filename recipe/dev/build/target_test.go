package build

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestTarget_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Target{})
}
