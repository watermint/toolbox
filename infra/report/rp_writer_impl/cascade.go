package rp_writer_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
)

func NewCascade(name string, c app_control.Control) rp_writer.Writer {
	fileWriters := make([]rp_writer.Writer, 0)
	consoleWriters := make([]rp_writer.Writer, 0)
	feature := c.Feature()

	// file writers
	switch {
	case feature.IsTransient():
		// no file writer

	case feature.BudgetMemory() == app_opt.BudgetLow:
		fileWriters = append(fileWriters, NewJsonWriter(name, c, false))

	default:
		fileWriters = append(fileWriters, NewJsonWriter(name, c, false))
		columnWriters := make([]rp_writer.Writer, 0)
		columnWriters = append(columnWriters, NewCsvWriter(name, c))
		columnWriters = append(columnWriters, NewXlsxWriter(name, c))
		sortedWriter := NewSorted(name, columnWriters)
		fileWriters = append(fileWriters, sortedWriter)
	}

	// console writer
	if !feature.IsQuiet() {
		if feature.UIFormat() == app_opt.OutputJson {
			consoleWriters = append(consoleWriters, NewJsonWriter(name, c, true))
		} else {
			consoleWriters = append(consoleWriters, newUIWriter(name, c))
		}
	}

	return &cascadeWriter{
		ctl:            c,
		name:           name,
		fileWriters:    fileWriters,
		consoleWriters: consoleWriters,
	}
}

type cascadeWriter struct {
	ctl            app_control.Control
	name           string
	writers        []rp_writer.Writer
	fileWriters    []rp_writer.Writer
	consoleWriters []rp_writer.Writer
	isClosed       bool
}

func (z *cascadeWriter) Name() string {
	return z.name
}

func (z *cascadeWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}

	z.writers = make([]rp_writer.Writer, 0)
	z.writers = append(z.writers, z.fileWriters...)
	if !ro.NoConsoleOutput {
		z.writers = append(z.writers, z.consoleWriters...)
	}

	for _, w := range z.writers {
		if err := w.Open(ctl, model, opts...); err != nil {
			z.Close()
			return err
		}
	}
	return nil
}

func (z *cascadeWriter) Row(r interface{}) {
	if z.isClosed {
		return
	}

	for _, w := range z.writers {
		w.Row(r)
	}
}

func (z *cascadeWriter) Close() {
	for _, w := range z.writers {
		w.Close()
	}
	z.isClosed = true
}
