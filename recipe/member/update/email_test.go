package update

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestEmail_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Email{})
}
