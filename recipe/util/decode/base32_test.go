package decode

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestBase32_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Base32{})
}
