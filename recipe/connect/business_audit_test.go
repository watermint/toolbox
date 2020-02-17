package connect

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestBusinessAudit_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &BusinessAudit{})
}
