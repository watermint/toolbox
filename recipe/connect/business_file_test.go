package connect

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestBusinessFile_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &BusinessFile{})
}
