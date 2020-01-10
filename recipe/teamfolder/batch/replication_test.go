package batch

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestReplication_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Replication{})
}
