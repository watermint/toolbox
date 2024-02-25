package mount

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMountable_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Mountable{})
}
