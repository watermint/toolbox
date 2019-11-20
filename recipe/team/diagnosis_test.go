package team

import (
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDiagnosis_Exec(t *testing.T) {
	qt_recipe.TestRecipe(t, &Diagnosis{})
}
