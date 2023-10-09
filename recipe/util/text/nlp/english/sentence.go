package english

import (
	"encoding/json"
	"github.com/jdkato/prose/v2"
	"github.com/watermint/toolbox/essentials/nlp/el_text"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Sentence struct {
	In              da_text.TextInput
	IgnoreLineBreak bool
}

func (z *Sentence) Preset() {
}

func (z *Sentence) Exec(c app_control.Control) error {
	content, err := z.In.Content()
	if err != nil {
		return err
	}

	inContent := el_text.IgnoreLineBreak(content, z.IgnoreLineBreak)

	doc, err := prose.NewDocument(inContent)
	if err != nil {
		return err
	}

	sentences := make([]string, 0)
	for _, s := range doc.Sentences() {
		sentences = append(sentences, s.Text)
	}
	out, err := json.Marshal(sentences)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, string(out))
	return nil
}

func (z *Sentence) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("english", "watermint toolbox is the best tool. I love watermint toolbox.")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Sentence{}, func(r rc_recipe.Recipe) {
		m := r.(*Sentence)
		m.In.SetFilePath(f)
	})
}
