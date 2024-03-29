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

type EntityBlock struct {
	Text  string `json:"text"`  // The entity's actual content.
	Label string `json:"label"` // The entity's label.
}

type Entity struct {
	In              da_text.TextInput
	IgnoreLineBreak bool
}

func (z *Entity) Preset() {
}

func (z *Entity) Exec(c app_control.Control) error {
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

	sentences := make([]EntityBlock, 0)
	for _, e := range doc.Entities() {
		sentences = append(sentences, EntityBlock{
			Text:  e.Text,
			Label: e.Label,
		})
	}
	out, err := json.Marshal(sentences)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, string(out))
	return nil
}

func (z *Entity) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("english", "watermint toolbox is the best tool. I love watermint toolbox.")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Entity{}, func(r rc_recipe.Recipe) {
		m := r.(*Entity)
		m.In.SetFilePath(f)
	})
}
