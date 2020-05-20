package es_generate

import (
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestStructTypeGenerator_Generate(t *testing.T) {
	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		t.Error(err)
		return
	}
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
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
		src, err := gen.Generate("source_generate_struct.go.tmpl")
		if err != nil {
			t.Error(err)
			return
		}
		out := es_stdout.NewDefaultOut(ctl.Feature())
		_, _ = out.Write(src)
	})
}
