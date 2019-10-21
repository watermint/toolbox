package activity

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestActivity_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Event{})
}
