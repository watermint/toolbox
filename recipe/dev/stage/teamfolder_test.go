package stage

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestTeamfolder_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Teamfolder{})
}
