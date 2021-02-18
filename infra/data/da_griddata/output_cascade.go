package da_griddata

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_opt"
)

func NewCascadeWriter(name, outputType, filePath string, formatters map[string]GridDataFormatter) GridDataWriter {
	return &cascadeWriter{
		name:       name,
		filePath:   filePath,
		outputType: outputType,
		formatters: formatters,
	}
}

type cascadeWriter struct {
	name       string
	filePath   string
	outputType string
	writers    []GridDataWriter
	formatters map[string]GridDataFormatter
}

func (z *cascadeWriter) Name() string {
	return z.name
}

func (z *cascadeWriter) Row(column []interface{}) {
	for _, w := range z.writers {
		w.Row(column)
	}
}

func (z *cascadeWriter) selectFormatter(outputType string) GridDataFormatter {
	if z.formatters == nil {
		return DefaultGridDataFormatter
	}
	if f, ok := z.formatters[outputType]; ok {
		return f
	}
	return DefaultGridDataFormatter
}

func (z *cascadeWriter) Open(c app_control.Control) error {
	z.writers = make([]GridDataWriter, 0)
	if !c.Feature().IsQuiet() {
		switch c.Feature().UIFormat() {
		case app_opt.OutputJson:
			z.writers = append(z.writers, NewConsoleWriter(z.selectFormatter(OutputTypeJson), NewJsonWriter()))
		default:
			z.writers = append(z.writers, NewConsoleWriter(z.selectFormatter(OutputTypeCsv), NewCsvWriter()))
		}
	}
	switch z.outputType {
	case OutputTypeAll:
		z.writers = append(z.writers, NewPlainWriter(z.name, z.filePath, z.selectFormatter(OutputTypeJson), NewJsonWriter()))
		z.writers = append(z.writers, NewPlainWriter(z.name, z.filePath, z.selectFormatter(OutputTypeCsv), NewCsvWriter()))

	case OutputTypeJson:
		z.writers = append(z.writers, NewPlainWriter(z.name, z.filePath, z.selectFormatter(OutputTypeJson), NewJsonWriter()))

	case OutputTypeCsv:
		z.writers = append(z.writers, NewPlainWriter(z.name, z.filePath, z.selectFormatter(OutputTypeCsv), NewCsvWriter()))
	}

	for _, w := range z.writers {
		if err := w.Open(c); err != nil {
			return err
		}
	}
	return nil
}

func (z *cascadeWriter) Close() {
	for _, w := range z.writers {
		w.Close()
	}
}
