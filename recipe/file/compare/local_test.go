package compare

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestLocal_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Local{})
}