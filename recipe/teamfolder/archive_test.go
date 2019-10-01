package teamfolder

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestArchive_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Archive{})
}
