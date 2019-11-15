package dev

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestDoc_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Doc{})
}
