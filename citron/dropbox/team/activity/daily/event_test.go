package daily

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestEvent_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Event{})
}
