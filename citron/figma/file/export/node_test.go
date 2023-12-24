package export

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestNode_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Node{})
}
