package connect

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUserFile_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &UserFile{})
}
