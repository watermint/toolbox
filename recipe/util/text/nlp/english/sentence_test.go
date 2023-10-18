package english

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestSentence_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Sentence{})
}
