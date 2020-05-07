package es_generate

import (
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"reflect"
	"testing"
)

func TestScannerImpl(t *testing.T) {
	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		t.Error(err)
		return
	}
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		sc, err := NewScanner(ctl, rr)
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
