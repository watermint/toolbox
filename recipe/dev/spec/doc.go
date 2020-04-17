package spec

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_doc"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"io"
	"os"
)

type Doc struct {
	rc_recipe.RemarkSecret
	Lang     mo_string.OptionalString
	FilePath mo_string.OptionalString
}

func (z *Doc) Preset() {
}

func (z *Doc) traverseCatalogue(c app_control.Control, cat rc_catalogue.Catalogue) error {
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
	if !z.FilePath.IsExists() {
		w = ut_io.NewDefaultOut(c.Feature().IsTest())
	} else {
		w, err = os.Create(z.FilePath.Value())
		if err != nil {
			l.Error("Unable to create spec file", zap.Error(err), zap.String("path", z.FilePath.Value()))
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

func (z *Doc) Exec(c app_control.Control) error {
	l := c.Log()
	if z.Lang.IsExists() {
		if c0, ok := app_control_launcher.ControlWithLang(z.Lang.Value(), c); ok {
			c = c0
		}
	}
	if cl, ok := c.(app_control_launcher.ControlLauncher); ok {
		return z.traverseCatalogue(c, cl.Catalogue())
	}
	l.Error("Not enough resource")
	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
