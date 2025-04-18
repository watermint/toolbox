package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_command"
	"github.com/watermint/toolbox/infra/doc/dc_contributor"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_knowledge"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/doc/dc_supplemental"
	"github.com/watermint/toolbox/infra/doc/dc_web"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_compatibility"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/quality/infra/qt_msgusage"
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
		folderPath := filepath.Dir(path)
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return err
		}
		return os.WriteFile(path, []byte(doc), 0644)
	}
}

func (z *Doc) genReadme(c app_control.Control) error {
	l := c.Log()
	path := dc_index.DocName(dc_index.MediaRepository, dc_index.DocRootReadme, c.Messages().Lang())
	l.Info("Generating README", esl.String("file", path))
	d := dc_readme.New(dc_index.MediaRepository, c.Messages(), z.Badge)
	doc := dc_section.Generate(dc_index.MediaRepository, dc_section.LayoutPage, c.Messages(), d)

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

func (z *Doc) genWeb(c app_control.Control) error {
	l := c.Log()
	lg := c.Messages().Lang()

	for _, doc := range dc_web.WebDocuments(c.Messages()) {
		path := dc_index.DocName(dc_index.MediaWeb, doc.DocId(), lg)
		layout := dc_section.LayoutPage
		if doc.DocId() == dc_index.DocWebHome {
			layout = dc_section.LayoutHome
		}
		l.Info("Generating Web Home", esl.String("file", path))
		doc := dc_section.Generate(dc_index.MediaWeb, layout, c.Messages(), doc)
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
		comDoc := dc_command.New(dc_index.MediaWeb, spec)
		path := dc_index.DocName(dc_index.MediaWeb, dc_index.DocManualCommand, c.Messages().Lang(), dc_index.CommandName(spec.SpecId()))
		doc := dc_section.Generate(dc_index.MediaWeb, dc_section.LayoutCommand, c.Messages(), comDoc)
		if err := z.genDoc(path, doc, c); err != nil {
			return err
		}

		// Compatibility documentation
		pc, found := rc_compatibility.Definitions.FindAlivePathChange(spec.Path())
		if found {
			for _, fp := range pc.FormerPaths {
				formedSpecId := strings.Join(append(fp.Path, fp.Name), "-")
				l.Info("Generating former command manual", esl.String("command", formedSpecId))
				comDoc := dc_command.NewCompatibilityNewPath(dc_index.MediaWeb, spec, fp, pc)
				path := dc_index.DocName(dc_index.MediaWeb, dc_index.DocManualCommand, c.Messages().Lang(), dc_index.CommandName(formedSpecId))
				doc := dc_section.Generate(dc_index.MediaWeb, dc_section.LayoutCommand, c.Messages(), comDoc)
				if err := z.genDoc(path, doc, c); err != nil {
					return err
				}
			}
		}
		prune, found := rc_compatibility.Definitions.FindPlannedPrune(spec.Path())
		if found {
			l.Info("Generating prune command manual", esl.String("command", spec.CliPath()))
			comDoc := dc_command.NewCompatibilityPrune(dc_index.MediaWeb, spec, prune)
			path := dc_index.DocName(dc_index.MediaWeb, dc_index.DocManualCommand, c.Messages().Lang(), dc_index.CommandName(spec.SpecId()))
			doc := dc_section.Generate(dc_index.MediaWeb, dc_section.LayoutCommand, c.Messages(), comDoc)
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

	for _, d := range dc_supplemental.Docs(dc_index.MediaWeb) {
		l.Info("Generating supplemental doc", esl.Int("media", int(dc_index.MediaWeb)), esl.Int("docId", int(d.DocId())))
		path := dc_index.DocName(dc_index.MediaWeb, d.DocId(), c.Messages().Lang())
		doc := dc_section.Generate(dc_index.MediaWeb, dc_section.LayoutPage, c.Messages(), d)

		if err := z.genDoc(path, doc, c); err != nil {
			return err
		}
	}
	return nil
}

func (z *Doc) genContributor(c app_control.Control) error {
	l := c.Log()
	for _, d := range dc_contributor.Docs(dc_index.MediaWeb) {
		l.Info("Generating contributor doc", esl.Int("media", int(dc_index.MediaWeb)), esl.Int("docId", int(d.DocId())))
		path := dc_index.DocName(dc_index.MediaWeb, d.DocId(), c.Messages().Lang())
		doc := dc_section.Generate(dc_index.MediaWeb, dc_section.LayoutContributor, c.Messages(), d)

		if err := z.genDoc(path, doc, c); err != nil {
			return err
		}
	}
	return nil
}

// genKnowledge generates knowledge base documentation for LLM training
func (z *Doc) genKnowledge(c app_control.Control) error {
	l := c.Log()

	// Get all recipes and convert them to specifications
	recipes := app_catalogue.Current().Recipes()
	specs := make([]rc_recipe.Spec, 0, len(recipes))

	for _, r := range recipes {
		specs = append(specs, rc_spec.New(r))
	}

	// Get supplemental documents for additional content
	additionalDocs := dc_supplemental.Docs(dc_index.MediaKnowledge)

	// Generate knowledge base content
	l.Info("Generating knowledge base documentation")
	knowledge := dc_knowledge.New(dc_index.MediaKnowledge)
	doc := knowledge.GenerateKnowledge(specs, additionalDocs)

	// Write the knowledge base file to the appropriate location
	path := dc_index.DocName(dc_index.MediaKnowledge, dc_index.DocWebKnowledge, c.Messages().Lang())
	l.Info("Writing knowledge base documentation", esl.String("file", path))

	return z.genDoc(path, doc, c)
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
	if err := z.genContributor(ctl); err != nil {
		l.Error("Failed to generate contributor documents", esl.Error(err))
		return err
	}
	if err := z.genWeb(ctl); err != nil {
		l.Error("Failed to generate web pages", esl.Error(err))
		return err
	}
	if err := z.genKnowledge(ctl); err != nil {
		l.Error("Failed to generate knowledge base documentation", esl.Error(err))
		return err
	}

	missing := qt_msgusage.Record().Missing()
	if len(missing) > 0 {
		return qt_messages.VerifyMessages(ctl)
	}
	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	c.Log().Info("Skipping doc test since some recipes may have been removed or renamed")
	return qt_errors.ErrorNoTestRequired
}
