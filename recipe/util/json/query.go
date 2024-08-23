package json

import (
	"encoding/json"
	"errors"
	"github.com/itchyny/gojq"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Query struct {
	rc_recipe.RemarkTransient
	Query   string `json:"query"`
	Path    da_text.TextInput
	Compact bool
}

func (z *Query) Preset() {
}

func (z *Query) Exec(c app_control.Control) error {
	content, err := z.Path.Content()
	if err != nil {
		return err
	}
	obj, err := es_json.ParseAny(content)
	if err != nil {
		c.Log().Warn("Unable to parse JSON", esl.Error(err), esl.String("content", string(content)))
		return err
	}
	q, err := gojq.Parse(z.Query)
	if err != nil {
		return err
	}
	iter := q.Run(obj)
	for {
		v, ok := iter.Next()
		if !ok {
			return nil
		}
		if err, ok := v.(error); ok {
			var hErr *gojq.HaltError
			if errors.As(err, &hErr) && hErr.Value() == nil {
				return nil
			}
			c.Log().Warn("Query error", esl.Error(err))
			return err
		}

		if z.Compact {
			v1, err := json.Marshal(v)
			if err != nil {
				return err
			}
			ui_out.TextOut(c, string(v1))
		} else {
			v1, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return err
			}
			ui_out.TextOut(c, string(v1))
		}
	}
}

func (z *Query) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("test.json", `{"a": 1, "b": 2}`)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Query{}, func(r rc_recipe.Recipe) {
		m := r.(*Query)
		m.Query = "."
		m.Path.SetFilePath(f)
	})
}
