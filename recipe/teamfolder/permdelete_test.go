package teamfolder

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestPermDelete_Exec(t *testing.T) {
	app_test.TestRecipe(t, &PermDelete{})
}
