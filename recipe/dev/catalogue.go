package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_source"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"reflect"
)

type Catalogue struct {
}

func (z *Catalogue) Preset() {
}

func (z *Catalogue) generateRecipe(rr string, sc ut_source.Scanner, c app_control.Control) error {
	l := c.Log()
	rcs := []string{"recipe", "ingredient"}
	for _, rc := range rcs {
		scr := sc.PathFilterPrefix(rc).ExcludeTest()
		sts, err := scr.FindStructImplements(reflect.TypeOf((*rc_recipe.Recipe)(nil)).Elem())
		if err != nil {
			return err
		}
		sg := ut_source.NewStructTypeGenerator(c, sts)
		op := filepath.Join(rr, "catalogue", rc+".go")
		f, err := os.Create(op)
		if err != nil {
			return err
		}

		l.Info("Generating recipe", zap.String("source", op))
		tmplName := fmt.Sprintf("catalogue_%s.go.tmpl", rc)
		if err := sg.Generate(tmplName, f); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

func (z *Catalogue) generateMessages(rr string, sc ut_source.Scanner, c app_control.Control) error {
	l := c.Log()
	scr := sc.ExcludeTest()
	sts, err := scr.FindStructHasPrefix("Msg")
	if err != nil {
		return err
	}
	sg := ut_source.NewStructTypeGenerator(c, sts)
	op := filepath.Join(rr, "catalogue/message.go")
	f, err := os.Create(op)
	if err != nil {
		return err
	}

	l.Info("Generating message", zap.String("source", op))
	tmplName := fmt.Sprintf("catalogue_message.go.tmpl")
	if err := sg.Generate(tmplName, f); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return nil
}

func (z *Catalogue) generateFeatures(rr string, sc ut_source.Scanner, c app_control.Control) error {
	l := c.Log()
	scr := sc.ExcludeTest()
	sts, err := scr.FindStructHasPrefix("OptInFeature")
	if err != nil {
		return err
	}
	sg := ut_source.NewStructTypeGenerator(c, sts)
	op := filepath.Join(rr, "catalogue/feature.go")
	f, err := os.Create(op)
	if err != nil {
		return err
	}

	l.Info("Generating feature", zap.String("source", op))
	tmplName := fmt.Sprintf("catalogue_feature.go.tmpl")
	if err := sg.Generate(tmplName, f); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return nil
}

func (z *Catalogue) Exec(c app_control.Control) error {
	rr, err := ut_source.DetectRepositoryRoot()
	if err != nil {
		return err
	}
	sc, err := ut_source.NewScanner(c, rr)
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
