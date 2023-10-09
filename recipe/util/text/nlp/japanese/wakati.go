package japanese

import (
	"github.com/watermint/toolbox/essentials/cache/ec_file"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/nlp/el_ja"
	"github.com/watermint/toolbox/essentials/nlp/el_text"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"strings"
)

type Wakati struct {
	In              da_text.TextInput
	Dictionary      mo_string.SelectString
	IgnoreLineBreak bool
	OmitBosEos      bool
	Separator       string
}

func (z *Wakati) Preset() {
	z.Dictionary.SetOptions("ipa", "ipa", "uni")
	z.Separator = " "
}

func (z *Wakati) Exec(c app_control.Control) error {
	content, err := z.In.Content()
	if err != nil {
		return err
	}

	inContent := el_text.IgnoreLineBreak(content, z.IgnoreLineBreak)

	cache := ec_file.New(c.Workspace().Cache(), c.Log())
	dc := el_ja.NewContainer(cache, c.Log())

	tok, err := dc.NewTokenizer(z.Dictionary.Value(), z.OmitBosEos)
	if err != nil {
		return err
	}

	segments := tok.Wakati(inContent)
	ui_out.TextOut(c, strings.Join(segments, z.Separator))
	return nil
}

func (z *Wakati) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("ja", "すもももももももものうち")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Wakati{}, func(r rc_recipe.Recipe) {
		m := r.(*Wakati)
		m.In.SetFilePath(f)
	})
}
