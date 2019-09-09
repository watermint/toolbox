package member

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestList_Exec(t *testing.T) {
	app_test.TestRecipe(t, &List{})
}
