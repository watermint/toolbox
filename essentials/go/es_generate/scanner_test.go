package es_generate

import (
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

func TestScannerImpl(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}
	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		t.Error(err)
		return
	}
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		sc, err := NewScanner(ctl, rr, ImporterTypeDefault)
		if err != nil {
			t.Error(err)
			return
		}
		sc = sc.ExcludeTest().PathFilterPrefix("recipe")

		sts, err := sc.FindStructHasPrefix("Msg")
		if err != nil {
			t.Error(err)
			return
		}

		if len(sts) < 1 {
			t.Error("not found")
		}

		sts, err = sc.FindStructImplements(reflect.TypeOf((*rc_recipe.Recipe)(nil)).Elem())
		if err != nil {
			t.Error(err)
		}
		if len(sts) < 1 {
			t.Error("not found")
		}
	})
}
