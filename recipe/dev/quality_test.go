package dev

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestQuality_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Quality{})
}
