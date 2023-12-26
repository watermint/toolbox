package account

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFilesystem_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Filesystem{})
}
