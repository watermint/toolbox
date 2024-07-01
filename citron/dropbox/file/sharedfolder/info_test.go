package sharedfolder

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestSharedFolder_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Info{})
}
