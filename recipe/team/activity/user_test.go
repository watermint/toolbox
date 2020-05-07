package activity

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUser_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &User{})
}
