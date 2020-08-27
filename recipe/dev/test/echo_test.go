package test

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestEcho_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Echo{})
}
