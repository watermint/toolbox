package build

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_command"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/doc/dc_supplemental"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/quality/infra/qt_msgusage"
	"io/ioutil"
)

type Doc struct {
	rc_recipe.RemarkSecret
	Badge   bool
	DocLang mo_string.OptionalString
}

func (z *Doc) Preset() {
	z.Badge = true
}

func (z *Doc) genDoc(path string, doc string, c app_control.Control) error {
	if c.Feature().IsTest() {
		out := es_stdout.NewDefaultOut(c.Feature())
		_, _ = fmt.Fprintln(out, doc)
		return nil
	} else {
		return ioutil.WriteFile(path, []byte(doc), 0644)
	}
}

func (z *Doc) genReadme(c app_control.Control) error {
	l := c.Log()
	path := dc_index.DocName(dc_index.MediaRepository, dc_index.DocRootReadme, c.Messages().Lang())
	l.Info("Generating README", esl.String("file", path))
	sec := dc_readme.New(dc_index.MediaRepository, c.Messages(), z.Badge)
	doc := dc_section.Generate(dc_index.MediaRepository, dc_section.LayoutPage, c.Messages(), sec...)

	return z.genDoc(path, doc, c)
}

func (z *Doc) genSecurity(c app_control.Control) error {
	l := c.Log()
	sec := dc_readme.NewSecurity()
	for _, m := range dc_index.AllMedia {
		path := dc_index.DocName(m, dc_index.DocRootSecurityAndPrivacy, c.Messages().Lang())
		l.Info("Generating SECURITY_AND_PRIVACY", esl.String("file", path))
		doc := dc_section.Generate(m, dc_section.LayoutPage, c.Messages(), sec)
		if err := z.genDoc(path, doc, c); err != nil {
			return err
		}
	}
	return nil
}

func (z *Doc) genCommands(c app_control.Control) error {
	recipes := app_catalogue.Current().Recipes()
	l := c.Log()

	for _, r := range recipes {
		spec := rc_spec.New(r)

		l.Info("Generating command manual", esl.String("command", spec.CliPath()))
		for _, m := range dc_index.AllMedia {
			sec := dc_command.New(m, spec)
			path := dc_index.DocName(m, dc_index.DocManualCommand, c.Messages().Lang(), dc_index.CommandName(spec.SpecId()))
			doc := dc_section.Generate(m, dc_section.LayoutCommand, c.Messages(), sec...)
			if err := z.genDoc(path, doc, c); err != nil {
				return err
			}
		}

		if err := qt_messages.SuggestCliArgs(c, r); err != nil {
			return err
		}
	}
	return nil
}

func (z *Doc) genSupplemental(c app_control.Control) error {
	l := c.Log()
	defer func() {
		if e := recover(); e != nil {
			switch re := e.(type) {
			case *rc_catalogue.RecipeNotFound:
				if c.Feature().IsTest() {
					l.Warn("Ignore recipe not found on test", esl.Error(re))
				} else {
					// re-throw
					panic(re)
				}
			default:
				// re-throw
				panic(re)
			}
		}
	}()
	for _, m := range dc_index.AllMedia {
		for _, d := range dc_supplemental.Docs(m) {
			l.Info("Generating supplemental doc", esl.Int("media", int(m)), esl.Int("docId", int(d.DocId())))
			path := dc_index.DocName(m, d.DocId(), c.Messages().Lang())
			doc := dc_section.Generate(m, dc_section.LayoutPage, c.Messages(), d.Sections()...)

			if err := z.genDoc(path, doc, c); err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *Doc) Exec(ctl app_control.Control) error {
	l := ctl.Log()

	if z.DocLang.IsExists() {
		ctl = ctl.WithLang(z.DocLang.Value())
	}
	if err := z.genReadme(ctl); err != nil {
		l.Error("Failed to generate README", esl.Error(err))
		return err
	}
	if err := z.genSecurity(ctl); err != nil {
		l.Error("Failed to generate SECURITY_AND_PRIVACY", esl.Error(err))
		return err
	}
	if err := z.genCommands(ctl); err != nil {
		l.Error("Failed to generate command manuals", esl.Error(err))
		return err
	}
	if err := z.genSupplemental(ctl); err != nil {
		l.Error("Failed to generate supplemental manuals", esl.Error(err))
		return err
	}

	missing := qt_msgusage.Record().Missing()
	if len(missing) > 0 {
		return qt_messages.VerifyMessages(ctl)
	}
	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
		rr := r.(*Doc)
		rr.Badge = false
	})
}
