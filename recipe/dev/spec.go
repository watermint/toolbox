package dev

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_doc"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"io"
	"os"
)

type Spec struct {
	Lang     string
	FilePath string
}

func (z *Spec) Preset() {
}

func (z *Spec) traverseCatalogue(c app_control.Control, cat rc_catalogue.Catalogue) error {
	l := c.Log()
	sd := make(map[string]*rc_doc.Recipe)

	for _, r := range cat.Recipes() {
		s := rc_spec.New(r)

		l.Debug("Generating", zap.String("recipe", s.CliPath()))
		d := s.Doc(c.UI())
		sd[d.Path] = d
	}

	var w io.WriteCloser
	var err error
	shouldClose := false
	if z.FilePath == "" {
		w = os.Stdout
	} else {
		w, err = os.Create(z.FilePath)
		if err != nil {
			l.Error("Unable to create spec file", zap.Error(err), zap.String("path", z.FilePath))
			return err
		}
		shouldClose = true
	}
	defer func() {
		if shouldClose {
			w.Close()
		}
	}()

	je := json.NewEncoder(w)
	je.SetIndent("", "  ")
	je.SetEscapeHTML(false)
	if err := je.Encode(sd); err != nil {
		l.Error("Unable to generate spec doc", zap.Error(err))
		return err
	}
	return nil
}

func (z *Spec) Exec(c app_control.Control) error {
	l := c.Log()
	if z.Lang != "" {
		wc := c.(app_control_launcher.WithMessageContainer)
		langPriority := make([]language.Tag, 0)
		ul := app_lang.Select(z.Lang)
		if ul != language.English {
			langPriority = append(langPriority, ul)
		}
		langPriority = append(langPriority, language.English)
		langContainers := make(map[language.Tag]app_msg_container.Container)

		for _, lang := range langPriority {
			mc, err := app_msg_container_impl.New(lang, c)
			if err != nil {
				return err
			}
			langContainers[lang] = mc
		}

		c = wc.With(app_msg_container_impl.NewMultilingual(langPriority, langContainers))
	}

	if cl, ok := c.(app_control_launcher.ControlLauncher); ok {
		return z.traverseCatalogue(c, cl.Catalogue())
	}
	l.Debug("Not enough resource")
	return nil
}

func (z *Spec) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
