package preflight

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestUp_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Up{})
}
