package update

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestExpiry_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Expiry{})
}
