package da_griddata

import (
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_unicode"
	"github.com/watermint/toolbox/essentials/io/es_file_read"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"io"
	"os"
)

type GridDataInput interface {
	SetFilePath(filePath string)
	FilePath() string
	EachRow(f func(col []interface{}, rowIndex int) error) error
	Open(ctl app_control.Control) error
	Spec() GridDataInputSpec
	Debug() interface{}
}

type MsgGridDataInput struct {
	ErrorUnableToRead app_msg.Message
}

var (
	MGridDataInput = app_msg.Apply(&MsgGridDataInput{}).(*MsgGridDataInput)
)

func NewInput(recipe interface{}, name string) GridDataInput {
	return &gdInput{
		recipe: recipe,
		name:   name,
	}
}

type gdInput struct {
	recipe   interface{}
	name     string
	filePath string
	ctl      app_control.Control
}

func (z *gdInput) Debug() interface{} {
	return map[string]interface{}{
		"Name":     z.name,
		"FilePath": z.filePath,
	}
}

func (z *gdInput) SetFilePath(filePath string) {
	z.filePath = filePath
}

func (z *gdInput) FilePath() string {
	return z.filePath
}

func (z *gdInput) EachRow(f func(col []interface{}, rowIndex int) error) error {
	ui := z.ctl.UI()
	l := z.ctl.Log().With(esl.String("path", z.filePath))
	rErr := es_file_read.ReadFileOrArchived(z.filePath, func(r io.Reader) error {
		cr := es_unicode.NewBomAwareCsvReader(r)
		for row := 0; ; row++ {
			ll := l.With(esl.Int("row", row))
			cols, err := cr.Read()
			switch err {
			case io.EOF:
				ll.Debug("Finished")
				return nil

			case nil:
				colData := make([]interface{}, len(cols))
				for i := range cols {
					colData[i] = cols[i]
				}
				if err = f(colData, row); err != nil {
					ll.Debug("Error returned from the operation", esl.Error(err))
					return err
				}

			default:
				ui.Error(MGridDataInput.ErrorUnableToRead.With("Error", err).With("Path", z.filePath))
				return err
			}
		}
	})
	if rErr != nil {
		ui.Error(MGridDataInput.ErrorUnableToRead.With("Error", rErr).With("Path", z.filePath))
		return rErr
	}
	return nil
}

func (z *gdInput) Open(ctl app_control.Control) error {
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

func (z *gdInput) Spec() GridDataInputSpec {
	return NewInputSpec(z.recipe, z.name)
}
