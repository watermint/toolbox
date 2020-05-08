package config

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestEnable_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Enable{})
}
