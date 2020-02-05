package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

func newCascade(name string, ctl app_control.Control) Writer {
	writers := make([]Writer, 0)

	writers = append(writers, NewJsonWriter(name, ctl, false))
	if !ctl.IsLowMemory() {
		writers = append(writers, newCsvWriter(name, ctl))
		writers = append(writers, NewXlsxWriter(name, ctl))
	}
	if ctl.IsQuiet() {
		writers = append(writers, NewJsonWriter(name, ctl, true))
	} else {
		writers = append(writers, newUIWriter(name, ctl))
	}

	return &cascadeWriter{
		ctl:     ctl,
		name:    name,
		writers: writers,
	}
}

type cascadeWriter struct {
	ctl      app_control.Control
	name     string
	writers  []Writer
	isClosed bool
}

func (z cascadeWriter) Name() string {
	return z.name
}

func (z cascadeWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
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

	p := z.ctl.Workspace().Report()
	ui := z.ctl.UI()
	ui.OpenArtifact(p, z.ctl.IsAutoOpen())
	z.isClosed = true
}
