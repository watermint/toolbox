package release

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAnnouncement_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Announcement{})
}
