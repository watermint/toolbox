package filerequest

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestClone_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Clone{})
}
