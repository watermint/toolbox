package team

import (
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"testing"
)

func TestDiagnosis_Exec(t *testing.T) {
	app_test.TestRecipe(t, &Diagnosis{})
}
