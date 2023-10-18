package english

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/cache/ec_file"
	"github.com/watermint/toolbox/essentials/nlp/el_en"
	"github.com/watermint/toolbox/essentials/nlp/el_text"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type TokenBlock struct {
	Tag   string `json:"tag"`   // The token's part-of-speech tag.
	Text  string `json:"text"`  // The token's actual content.
	Label string `json:"label"` // The token's IOB label.
}

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
	dc := el_en.NewContainer(cache, c.Log())
	doc, err := dc.NewDocument(inContent)

	if err != nil {
		return err
	}

	sentences := make([]TokenBlock, 0)
	for _, t := range doc.Tokens() {
		sentences = append(sentences, TokenBlock{
			Text:  t.Text,
			Label: t.Label,
		})
	}
	out, err := json.Marshal(sentences)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, string(out))
	return nil
}

func (z *Token) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("english", "watermint toolbox is the best tool. I love watermint toolbox.")
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
