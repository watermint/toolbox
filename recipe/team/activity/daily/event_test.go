package daily

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestEvent_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Event{})
}
