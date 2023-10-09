package japanese

import (
	"encoding/json"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/watermint/toolbox/essentials/cache/ec_file"
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
	IgnoreLineBreak bool
}

func (z *Token) Preset() {
}

func (z *Token) Exec(c app_control.Control) error {
	content, err := z.In.Content()
	if err != nil {
		return err
	}

	inContent := el_text.IgnoreLineBreak(content, z.IgnoreLineBreak)

	cache := ec_file.New(c.Workspace().Cache(), c.Log())
	dc := el_ja.NewContainer(cache, c.Log())
	dic, err := dc.Load("ipa")
	if err != nil {
		return err
	}
	tok, err := tokenizer.New(dic)
	if err != nil {
		return err
	}
	tokens := el_ja.NewTokenArray(tok.Tokenize(inContent))
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
