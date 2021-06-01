package message

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestSend_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Send{})
}
