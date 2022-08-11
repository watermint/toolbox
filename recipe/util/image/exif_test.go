package image

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExif_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Exif{})
}
