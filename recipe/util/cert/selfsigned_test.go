package cert

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestSelfsigned_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Selfsigned{})
}
