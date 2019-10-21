package _import

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestUrl_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Url{})
}
