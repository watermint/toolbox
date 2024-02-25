package sharedfolder

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestShare_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Share{})
}
