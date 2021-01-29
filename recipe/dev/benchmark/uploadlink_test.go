package benchmark

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUploadlink_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Uploadlink{})
}
