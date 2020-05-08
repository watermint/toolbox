package content

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMember_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Member{})
}
