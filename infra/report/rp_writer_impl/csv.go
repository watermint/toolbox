package rp_writer_impl

import (
	"encoding/csv"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_column_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"os"
	"path/filepath"
	"sync"
)

func NewCsvWriter(name string, ctl app_control.Control) rp_writer.Writer {
	return &csvWriter{
		name: name,
		ctl:  ctl,
	}
}

type csvWriter struct {
	name     string
	index    int
	path     string
	file     *os.File
	w        *csv.Writer
	mutex    sync.Mutex
	ctl      app_control.Control
	colModel rp_column.Column
}

func (z *csvWriter) Name() string {
	return z.name
}

func (z *csvWriter) Row(r interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.index == 0 {
		z.w.Write(z.colModel.Header())
	}
	z.w.Write(z.colModel.ValueStrings(r))
	z.index++
}

func (z *csvWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) (err error) {
	z.ctl = ctl
	l := ctl.Log()
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}

	z.colModel = rp_column_impl.NewModel(model, opts...)
	z.path = filepath.Join(ctl.Workspace().Report(), z.Name()+ro.ReportSuffix+".csv")
	l = l.With(esl.String("path", z.path))
	l.Debug("Create new csv report")
	z.file, err = os.Create(z.path)
	if err != nil {
		l.Error("Unable to create file", esl.Error(err))
		return err
	}
	z.w = csv.NewWriter(z.file)
	return nil
}

func (z *csvWriter) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.ctl.Log().With(esl.String("path", z.path))

	if z.file != nil {
		z.w.Flush()
		z.file.Sync()
		err := z.file.Close()
		l.Debug("File closed", esl.Error(err))

		if z.index < 1 && z.ctl.Feature().IsProduction() && !z.ctl.Feature().IsTest() {
			l.Debug("Try removing empty report file")
			err := os.Remove(z.path)
			l.Debug("Removed or had an error (ignore)", esl.Error(err))
		}
		z.file = nil
		z.w = nil
	}
}
