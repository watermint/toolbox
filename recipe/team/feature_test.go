package team

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestFeature_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Feature{})
}
