package da_griddata

import (
	"errors"
	"flag"
	"strings"

	"github.com/watermint/toolbox/essentials/encoding/es_json"
	es_case2 "github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewOutput(recipe interface{}, name string) GridDataOutput {
	return &gdOutput{
		recipe:     recipe,
		name:       name,
		formatters: make(map[string]GridDataFormatter),
		writer:     NewConsoleWriter(DefaultGridDataFormatter, NewCsvWriter()),
	}
}

var (
	ErrorValueRestoreFailed = errors.New("value restore failed")
)

const (
	GridOutputOptionFilePath = ""
	GridOutputOptionFormat   = "Format"
	GridOutputDescSuffix     = ".output_grid_data.desc"
)

type gdOutput struct {
	recipe     interface{}
	name       string
	filePath   string
	outputType string
	formatters map[string]GridDataFormatter
	writer     GridDataWriter
}

func (z *gdOutput) FilePath() string {
	return z.filePath
}

func (z *gdOutput) Debug() interface{} {
	return map[string]interface{}{
		"Name":       z.name,
		"FilePath":   z.filePath,
		"OutputType": z.outputType,
	}
}

func (z *gdOutput) Name() string {
	return z.name
}

func (z *gdOutput) ApplyFlags(f *flag.FlagSet, fieldDesc app_msg.Message, ui app_ui.UI) {
	// descFilePath := ui.Text(app_msg.CreateMessage(fieldDesc.Key() + es_case.ToLowerSnakeCase(GridOutputOptionFilePath)))
	// descFormat := ui.Text(app_msg.CreateMessage(fieldDesc.Key() + es_case.ToLowerSnakeCase(GridOutputOptionFormat)))
	descFilePath := ui.Text(z.FieldDesc(fieldDesc, z.Name()+GridOutputOptionFilePath))
	descFormat := ui.Text(z.FieldDesc(fieldDesc, z.Name()+GridOutputOptionFormat))
	f.StringVar(&z.filePath, es_case2.ToLowerKebabCase(z.Name()+GridOutputOptionFilePath), "", descFilePath)
	f.StringVar(&z.outputType, es_case2.ToLowerKebabCase(z.Name()+GridOutputOptionFormat), OutputTypeCsv, descFormat)
}

func (z *gdOutput) Fields() []string {
	return []string{
		z.Name() + GridOutputOptionFilePath,
		z.Name() + GridOutputOptionFormat,
	}
}

func (z *gdOutput) FieldDesc(base app_msg.Message, name string) app_msg.Message {
	fieldName := name
	if strings.HasPrefix(fieldName, z.Name()) {
		fieldName = name[len(z.Name()):]
	}
	if fieldName == "" {
		return app_msg.CreateMessage(base.Key() + GridOutputDescSuffix)
	} else {
		return app_msg.CreateMessage(base.Key() + "." + es_case2.ToLowerSnakeCase(fieldName) + ".output_grid_data")
	}
}

func (z *gdOutput) Capture() interface{} {
	s := make(map[string]string)
	s[GridOutputOptionFilePath] = z.filePath
	s[GridOutputOptionFormat] = z.outputType
	return s
}

func (z *gdOutput) Restore(v es_json.Json) error {
	if obj, found := v.Object(); found {
		if s, found := obj[GridOutputOptionFilePath]; found {
			if v, found := s.String(); found {
				z.filePath = v
			}
		}
		if s, found := obj[GridOutputOptionFormat]; found {
			if v, found := s.String(); found {
				z.outputType = v
			}
		}
		return nil
	}
	return ErrorValueRestoreFailed
}

func (z *gdOutput) Row(column []interface{}) {
	z.writer.Row(column)
}

func (z *gdOutput) SetFormatter(outType string, formatter GridDataFormatter) {
	z.formatters[outType] = formatter
}

func (z *gdOutput) Open(c app_control.Control) error {
	z.writer = NewCascadeWriter(z.Name(), z.outputType, z.filePath, z.formatters)
	return z.writer.Open(c)
}

func (z *gdOutput) Close() {
	z.writer.Close()
}

func (z *gdOutput) Spec() GridDataOutputSpec {
	return NewOutputSpec(z.recipe, z.name)
}
