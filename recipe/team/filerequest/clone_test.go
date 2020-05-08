package filerequest

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestClone_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Clone{})
}
