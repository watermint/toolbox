package ig_teamfolder

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestReplication_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Replication{})
}
