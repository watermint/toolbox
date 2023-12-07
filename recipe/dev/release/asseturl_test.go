package release

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAsseturl_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Asseturl{})
}
