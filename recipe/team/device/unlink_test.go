package device

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestUnlink_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Unlink{})
}
