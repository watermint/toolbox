package msg

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAdd_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Add{})
}