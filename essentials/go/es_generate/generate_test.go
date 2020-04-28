package es_generate

import (
	"github.com/watermint/toolbox/essentials/io/ut_io"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestStructTypeGenerator_Generate(t *testing.T) {
	rr, err := DetectRepositoryRoot()
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
		sc = sc.PathFilterPrefix("recipe")
		sts, err := sc.FindStructHasPrefix("Msg")
		if err != nil {
			t.Error(err)
			return
		}
		gen := NewStructTypeGenerator(ctl, sts)
		err = gen.Generate("source_generate_struct.go.tmpl", ut_io.NewDefaultOut(ctl.Feature().IsTest()))
		if err != nil {
			t.Error(err)
			return
		}
	})
}
