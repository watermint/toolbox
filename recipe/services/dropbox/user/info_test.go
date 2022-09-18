package user

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestInfo_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Info{})
}
