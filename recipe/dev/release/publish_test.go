package release

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPublish_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Publish{})
}
