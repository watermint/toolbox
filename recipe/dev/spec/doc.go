package spec

import (
	"compress/gzip"
	"encoding/json"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
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

func (z *Doc) traverseCatalogue(c app_control.Control) error {
	l := c.Log()
	sd := make(map[string]*dc_recipe.Recipe)
	cat := app_catalogue.Current()

	for _, r := range cat.Recipes() {
		s := rc_spec.New(r)

		l.Debug("Generating", esl.String("recipe", s.CliPath()))
		d := s.Doc(c.UI())
		sd[d.Path] = d
	}

	jeOut := func(w io.Writer) error {
		je := json.NewEncoder(w)
		je.SetIndent("", "  ")
		je.SetEscapeHTML(false)
		if err := je.Encode(sd); err != nil {
			l.Error("Unable to generate spec doc", esl.Error(err))
			return err
		}
		return nil
	}

	var w io.WriteCloser
	var err error
	if !z.FilePath.IsExists() {
		w = es_stdout.NewDefaultOut(c.Feature())
		return jeOut(w)
	}

	w, err = os.Create(z.FilePath.Value())
	if err != nil {
		l.Error("Unable to create spec file", esl.Error(err), esl.String("path", z.FilePath.Value()))
		return err
	}
	defer func() {
		_ = w.Close()
	}()

	gw := gzip.NewWriter(w)
	defer func() {
		_ = gw.Flush()
		_ = gw.Close()
	}()
	return jeOut(gw)
}

func (z *Doc) Exec(c app_control.Control) error {
	if z.Lang.IsExists() {
		return z.traverseCatalogue(c.WithLang(z.Lang.Value()))
	} else {
		return z.traverseCatalogue(c)
	}
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
