package member

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestInvite_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Invite{})
}
