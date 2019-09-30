package quota

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestUsage_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Usage{})
}
