package build

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_generate"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"os"
	"path/filepath"
	"reflect"
)

type Catalogue struct {
	rc_recipe.RemarkSecret
	Importer mo_string.SelectString
}

func (z *Catalogue) Preset() {
	z.Importer.SetOptions(
		string(es_generate.ImporterTypeDefault),
		string(es_generate.ImporterTypeDefault),
		string(es_generate.ImporterTypeEnhanced),
	)
}

func (z *Catalogue) generateRecipe(rr string, sc es_generate.Scanner, c app_control.Control) error {
	l := c.Log()
	rcs := app_definitions.RecipePackageNames
	for _, rc := range rcs {
		l.Debug("Scanning recipe", esl.String("path", rc))
		scr := sc.PathFilterPrefix(rc).ExcludeTest()
		sts, err := scr.FindStructImplements(reflect.TypeOf((*rc_recipe.Recipe)(nil)).Elem())
		if err != nil {
			return err
		}
		sg := es_generate.NewStructTypeGenerator(c, sts)
		op := filepath.Join(rr, "catalogue", rc+".go")

		l.Info("Generating recipe", esl.String("source", op))
		tmplName := fmt.Sprintf("catalogue_%s.go.tmpl", rc)
		src, err := sg.Generate(tmplName)
		if err != nil {
			l.Debug("Unable to generate", esl.Error(err))
			return err
		}
		err = os.WriteFile(op, src, 0644)
		if err != nil {
			l.Debug("Unable to write", esl.Error(err))
			return err
		}
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

	l.Info("Generating message", esl.String("source", op))
	tmplName := fmt.Sprintf("catalogue_message.go.tmpl")
	src, err := sg.Generate(tmplName)
	if err != nil {
		l.Debug("Unable to generate", esl.Error(err))
		return err
	}
	return os.WriteFile(op, src, 0644)
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

	l.Info("Generating feature", esl.String("source", op))
	tmplName := fmt.Sprintf("catalogue_feature.go.tmpl")
	src, err := sg.Generate(tmplName)
	if err != nil {
		return err
	}
	return os.WriteFile(op, src, 0644)
}

func (z *Catalogue) Exec(c app_control.Control) error {
	l := c.Log()
	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}
	l.Info("Scanning repository root", esl.String("path", rr))
	sc, err := es_generate.NewScanner(c, rr, es_generate.ImporterType(z.Importer.Value()))
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
