package util

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCurl_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Curl{})
}
