package ig_file

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDownload_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Download{})
}
