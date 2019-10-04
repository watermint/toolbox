package teamfolder

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestReplication_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Replication{})
}
