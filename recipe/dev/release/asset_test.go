package release

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAsset_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Asset{})
}
