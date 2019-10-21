package sharedlink

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestCreate_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Create{})
}
