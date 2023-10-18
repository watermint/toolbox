package japanese

import (
	"encoding/json"
	"github.com/ikawaha/kagome/v2/tokenizer"
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
)

type Token struct {
	In              da_text.TextInput
	Dictionary      mo_string.SelectString
	IgnoreLineBreak bool
	OmitBosEos      bool
	Mode            mo_string.SelectString
}

func (z *Token) Preset() {
	z.Dictionary.SetOptions(el_ja.DictionaryIpa, el_ja.DictionaryIpa, el_ja.DictionaryUni)
	z.Mode.SetOptions(el_ja.ModeNormal, el_ja.ModeNormal, el_ja.ModeSearch, el_ja.ModeExtend)
}

func (z *Token) Exec(c app_control.Control) error {
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

	var mode tokenizer.TokenizeMode
	switch z.Mode.Value() {
	case el_ja.ModeNormal:
		mode = tokenizer.Normal
	case el_ja.ModeSearch:
		mode = tokenizer.Search
	case el_ja.ModeExtend:
		mode = tokenizer.Extended
	default:
		mode = tokenizer.Normal
	}

	tokens := el_ja.NewTokenArray(tok.Analyze(inContent, mode))
	out, err := json.Marshal(tokens)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, string(out))
	return nil
}

func (z *Token) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("ja", "すもももももももものうち")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Token{}, func(r rc_recipe.Recipe) {
		m := r.(*Token)
		m.In.SetFilePath(f)
	})
}
