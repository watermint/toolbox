package decode

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestBase64_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Base64{})
}
