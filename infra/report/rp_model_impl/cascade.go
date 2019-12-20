package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

func newCascade(name string, ctl app_control.Control) Writer {
	return &cascadeWriter{
		name: name,
		writers: []Writer{
			&jsonWriter{name: name, ctl: ctl},
		},
	}
}

type cascadeWriter struct {
	name    string
	writers []Writer
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
	for _, w := range z.writers {
		w.Row(r)
	}
}

func (z *cascadeWriter) Close() {
	for _, w := range z.writers {
		w.Close()
	}
}
