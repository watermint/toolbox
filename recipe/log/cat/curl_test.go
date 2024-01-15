package cat

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCurl_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Curl{})
}
