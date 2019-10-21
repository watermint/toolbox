package compare

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestAccount_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Account{})
}
