package doc

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMarkdown_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Markdown{})
}
