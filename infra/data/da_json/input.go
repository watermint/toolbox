package da_json

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/io/es_file_read"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgJsonInput struct {
	WarnUnableParseTheLineSkip app_msg.Message
}

var (
	MJsonInput = app_msg.Apply(&MsgJsonInput{}).(*MsgJsonInput)
)

type JsonInput interface {
	SetFilePath(filePath string)
	FilePath() string
	Spec() JsonInputSpec
	Debug() interface{}
	Open(ctl app_control.Control) error

	// Unmarshal unmarshals the file content to the model.
	Unmarshal() (interface{}, error)

	SetModel(model interface{})

	EachModel(f func(m interface{}) error) error
}

func NewInput(name string, recipe interface{}) JsonInput {
	return &jsInput{
		name:   name,
		recipe: recipe,
	}
}

type jsInput struct {
	recipe   interface{}
	model    interface{}
	name     string
	filePath string
	ctl      app_control.Control
}

func (z *jsInput) Unmarshal() (v interface{}, err error) {
	fileContent, err := os.ReadFile(z.filePath)
	if err != nil {
		return nil, err
	}
	v = es_reflect.NewInstance(z.model)
	err = json.Unmarshal(fileContent, v)
	return
}

func (z *jsInput) SetFilePath(filePath string) {
	z.filePath = filePath
}

func (z *jsInput) FilePath() string {
	return z.filePath
}

func (z *jsInput) Spec() JsonInputSpec {
	return NewJsonSpec(z.name, z.recipe)
}

func (z *jsInput) Debug() interface{} {
	if z.model != nil {
		return map[string]interface{}{
			"Name":     z.name,
			"FilePath": z.filePath,
			"Model":    es_reflect.Key(z.model),
		}
	} else {
		return map[string]interface{}{
			"Name":           z.name,
			"FilePath":       z.filePath,
			"NoModelDefined": true,
		}
	}
}

func (z *jsInput) Open(ctl app_control.Control) error {
	l := ctl.Log().With(esl.String("path", z.filePath))
	z.ctl = ctl
	ls, err := os.Lstat(z.filePath)
	if err != nil {
		l.Debug("Unable to locate the file", esl.Error(err))
		return err
	}
	if ls.IsDir() {
		l.Debug("The path is not a file", esl.Any("lstat", ls))
		return errors.New("not a file")
	}
	l.Debug("Located Grid Data Input")
	return nil
}

func (z *jsInput) SetModel(model interface{}) {
	z.model = model
}

func (z *jsInput) EachModel(f func(m interface{}) error) error {
	ui := z.ctl.UI()
	l := z.ctl.Log().With(esl.String("filePath", z.filePath))
	l.Debug("Load file")
	rErr := es_file_read.ReadFileOrArchived(z.filePath, func(r io.Reader) error {
		data, err := io.ReadAll(r)
		if err != nil {
			l.Debug("Unable to read", esl.Error(err))
			return err
		}

		// Try to parse as JSON array.
		if j, err := es_json.Parse(data); err != nil {
			l.Debug("Unable to parse, fallback to JSON line parse", esl.Error(err))
		} else {
			err := j.ArrayEach(func(e es_json.Json) error {
				v := es_reflect.NewInstance(z.model)
				if jErr := e.Model(v); jErr != nil {
					l.Debug("Unable parse as the model", esl.Error(jErr))
					return jErr
				}
				return f(v)
			})

			if errors.Is(err, es_json.ErrorNotAnArray) {
				// Try to parse as a single object.
				v := es_reflect.NewInstance(z.model)
				if err := j.Model(v); err != nil {
					l.Debug("Unable to parse", esl.Error(err))
					return err
				}
				return f(v)
			}
			if err != nil {
				l.Debug("Error on process", esl.Error(err))
				return err
			}
			return nil
		}

		// Try to parse as JSON lines.

		scanner := bufio.NewScanner(bytes.NewReader(data))
		lines := 0
		for scanner.Scan() {
			line := scanner.Bytes()
			lines++
			if len(strings.TrimSpace(string(line))) < 1 {
				l.Debug("Skip empty line", esl.Int("line", lines))
				continue
			}
			v := es_reflect.NewInstance(z.model)
			j, err := es_json.Parse(line)
			if err != nil {
				l.Debug("Unable to parse the line", esl.Error(err))
				ui.Error(MJsonInput.WarnUnableParseTheLineSkip.With("Error", err).With("Line", lines))
				continue
			}
			if err := j.Model(v); err != nil {
				l.Debug("Unable to parse the model", esl.Error(err))
				ui.Error(MJsonInput.WarnUnableParseTheLineSkip.With("Error", err).With("Line", lines))
				continue
			}
			if err := f(v); err != nil {
				l.Debug("The function returned an error", esl.Error(err))
				return err
			}
		}
		if err := scanner.Err(); err != nil {
			l.Debug("Error during read", esl.Error(err))
			return err
		}
		l.Debug("Successfully scanned the text", esl.Int("lines", lines))
		return nil

	})
	return rErr
}
