package archive

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUnzip_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Unzip{})
}
