package test

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestAuth_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Auth{})
}
