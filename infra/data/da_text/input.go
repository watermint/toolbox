package da_text

import (
	"bufio"
	"errors"
	"github.com/watermint/toolbox/essentials/io/es_file_read"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"io"
	"os"
)

type TextInput interface {
	SetFilePath(filePath string)
	FilePath() string
	Spec() TextInputSpec
	Debug() interface{}
	Open(ctl app_control.Control) error

	// Read each line. If a func f returns an error, finish read text immediately and returns the error.
	EachLine(f func(line string) error) error

	// Returns entire text content.
	Content() ([]byte, error)
}

func NewTextInput(name string, recipe interface{}) TextInput {
	return &txInput{
		recipe:   recipe,
		name:     name,
		filePath: "",
		ctl:      nil,
	}
}

type txInput struct {
	recipe   interface{}
	name     string
	filePath string
	ctl      app_control.Control
}

func (z *txInput) SetFilePath(filePath string) {
	z.filePath = filePath
}

func (z *txInput) FilePath() string {
	return z.filePath
}

func (z *txInput) Spec() TextInputSpec {
	return NewInputSpec(z.name, z.recipe)
}

func (z *txInput) Debug() interface{} {
	return map[string]interface{}{
		"Name":     z.name,
		"FilePath": z.filePath,
	}
}

func (z *txInput) Open(ctl app_control.Control) error {
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

func (z *txInput) EachLine(f func(line string) error) error {
	l := z.ctl.Log().With(esl.String("filePath", z.filePath))
	l.Debug("Load file")
	rErr := es_file_read.ReadFileOrArchived(z.filePath, func(r io.Reader) error {
		scanner := bufio.NewScanner(r)
		lines := 0
		for scanner.Scan() {
			if err := f(scanner.Text()); err != nil {
				l.Debug("The function returned an error", esl.Error(err))
				return err
			}
			lines++
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

func (z *txInput) Content() (content []byte, err error) {
	l := z.ctl.Log().With(esl.String("filePath", z.filePath))

	// If the file path is "-", read from stdin
	if z.filePath == "-" {
		content, err = io.ReadAll(os.Stdin)
		l.Debug("Content load finished", esl.Int("contentLength", len(content)), esl.Error(err))
		return
	}

	content, err = os.ReadFile(z.filePath)
	l.Debug("Content load finished", esl.Int("contentLength", len(content)), esl.Error(err))
	return
}
