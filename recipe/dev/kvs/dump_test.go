package kvs

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDump_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Dump{})
}