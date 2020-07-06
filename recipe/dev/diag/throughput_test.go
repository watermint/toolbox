package diag

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestThroughput_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Throughput{})
}
