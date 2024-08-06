package uuid

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUlid_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Ulid{})
}
