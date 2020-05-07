package dev

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_generate"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"os"
	"path/filepath"
	"reflect"
)

type Catalogue struct {
	rc_recipe.RemarkSecret
}

func (z *Catalogue) Preset() {
}

func (z *Catalogue) generateRecipe(rr string, sc es_generate.Scanner, c app_control.Control) error {
	l := c.Log()
	rcs := []string{"recipe", "ingredient"}
	for _, rc := range rcs {
		scr := sc.PathFilterPrefix(rc).ExcludeTest()
		sts, err := scr.FindStructImplements(reflect.TypeOf((*rc_recipe.Recipe)(nil)).Elem())
		if err != nil {
			return err
		}
		sg := es_generate.NewStructTypeGenerator(c, sts)
		op := filepath.Join(rr, "catalogue", rc+".go")
		f, err := os.Create(op)
		if err != nil {
			return err
		}

		l.Info("Generating recipe", es_log.String("source", op))
		tmplName := fmt.Sprintf("catalogue_%s.go.tmpl", rc)
		if err := sg.Generate(tmplName, f); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

func (z *Catalogue) generateMessages(rr string, sc es_generate.Scanner, c app_control.Control) error {
	l := c.Log()
	scr := sc.ExcludeTest()
	sts, err := scr.FindStructHasPrefix("Msg")
	if err != nil {
		return err
	}
	sg := es_generate.NewStructTypeGenerator(c, sts)
	op := filepath.Join(rr, "catalogue/message.go")
	f, err := os.Create(op)
	if err != nil {
		return err
	}

	l.Info("Generating message", es_log.String("source", op))
	tmplName := fmt.Sprintf("catalogue_message.go.tmpl")
	if err := sg.Generate(tmplName, f); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return nil
}

func (z *Catalogue) generateFeatures(rr string, sc es_generate.Scanner, c app_control.Control) error {
	l := c.Log()
	scr := sc.ExcludeTest()
	sts, err := scr.FindStructHasPrefix("OptInFeature")
	if err != nil {
		return err
	}
	sg := es_generate.NewStructTypeGenerator(c, sts)
	op := filepath.Join(rr, "catalogue/feature.go")
	f, err := os.Create(op)
	if err != nil {
		return err
	}

	l.Info("Generating feature", es_log.String("source", op))
	tmplName := fmt.Sprintf("catalogue_feature.go.tmpl")
	if err := sg.Generate(tmplName, f); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return nil
}

func (z *Catalogue) Exec(c app_control.Control) error {
	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}
	sc, err := es_generate.NewScanner(c, rr)
	if err != nil {
		return err
	}
	if err := z.generateRecipe(rr, sc, c); err != nil {
		return err
	}
	if err := z.generateMessages(rr, sc, c); err != nil {
		return err
	}
	if err := z.generateFeatures(rr, sc, c); err != nil {
		return err
	}

	return nil
}

func (z *Catalogue) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Catalogue{}, rc_recipe.NoCustomValues)
}
