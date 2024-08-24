package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/itchyny/gojq"
	"github.com/simonfrey/jsonl"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io"
	"os"
)

type Query struct {
	rc_recipe.RemarkTransient
	Query   string `json:"query"`
	Path    da_text.TextInput
	Compact bool
	Lines   bool
}

func (z *Query) Preset() {
}

func (z *Query) execObject(c app_control.Control, obj interface{}) error {
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
			switch v1 := v.(type) {
			case string:
				ui_out.TextOut(c, v1)
			case int, float64, bool:
				ui_out.TextOut(c, fmt.Sprintf("%d", v1))
			default:
				v2, err := json.Marshal(v1)
				if err != nil {
					ui_out.TextOut(c, fmt.Sprintf("%v", v))
					continue
				}
				ui_out.TextOut(c, string(v2))
			}
		} else {
			switch v1 := v.(type) {
			case string:
				ui_out.TextOut(c, v1)
			case int, float64, bool:
				ui_out.TextOut(c, fmt.Sprintf("%d", v1))
			default:
				v2, err := json.MarshalIndent(v, "", "  ")
				if err != nil {
					ui_out.TextOut(c, fmt.Sprintf("%v", v))
					continue
				}
				ui_out.TextOut(c, string(v2))
			}
		}
	}
}

func (z *Query) execContent(c app_control.Control, content []byte) error {
	obj, err := es_json.ParseAny(content)
	if err != nil {
		c.Log().Warn("Unable to parse JSON", esl.Error(err), esl.String("content", string(content)))
		return err
	}

	return z.execObject(c, obj)
}

func (z *Query) execLines(c app_control.Control, path string) error {
	l := c.Log()
	var source io.ReadCloser
	if path == "-" {
		l.Debug("Read from stdin")
		source = os.Stdin
	} else {
		l.Debug("Read from file", esl.String("path", path))
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()
		source = f
	}
	jsl := jsonl.NewReader(source)
	lineNum := 0
	return jsl.ReadLines(func(data []byte) error {
		var obj interface{}
		if err := json.Unmarshal(data, &obj); err != nil {
			l.Debug("Unable to parse JSON", esl.Error(err), esl.String("line", string(data)))
			return err
		}
		l.Debug("Read", esl.Any("obj", obj), esl.Int("lineNum", lineNum))
		if err := z.execObject(c, obj); err != nil {
			return err
		}
		lineNum++
		return nil
	})
}

func (z *Query) Exec(c app_control.Control) error {
	if z.Lines {
		return z.execLines(c, z.Path.FilePath())
	} else {
		content, err := z.Path.Content()
		if err != nil {
			return err
		}
		return z.execContent(c, content)
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
