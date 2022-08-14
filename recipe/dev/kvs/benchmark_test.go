package kvs

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestBenchmark_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Benchmark{})
}
