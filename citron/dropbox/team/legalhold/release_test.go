package legalhold

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestRelease_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Release{})
}
