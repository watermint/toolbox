package rp_writer_impl

import (
	"bufio"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/text/es_sort"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_column_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"path/filepath"
	"sync"
)

type MsgSortedWriter struct {
	ProgressSorting          app_msg.Message
	ProgressPreparing        app_msg.Message
	ErrorUnableComposeReport app_msg.Message
}

var (
	MSortedWriter = app_msg.Apply(&MsgSortedWriter{}).(*MsgSortedWriter)
)

func NewSorted(name string, writers []rp_writer.Writer) rp_writer.Writer {
	return &Sorted{
		name:    name,
		writers: writers,
	}
}

type Sorted struct {
	ctl     app_control.Control
	name    string
	writers []rp_writer.Writer
	dstPath string
	dstFile *os.File
	sorter  es_sort.Sorter
	stream  rp_column.Column
	json    rp_column.Column
	isOpen  bool
	mutex   sync.Mutex
}

func (z *Sorted) Name() string {
	return z.name
}

func (z *Sorted) Row(r interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.ctl.Log().With(esl.String("name", z.name))

	if !z.isOpen {
		l.Warn("The writer is not yet open")
		return
	}

	vals := z.stream.Values(r)

	b, err := json.Marshal(vals)
	if err != nil {
		l.Warn("Unable to marshal", esl.Error(err))
		return
	}

	err = z.sorter.WriteLine(string(b))
	app_ui.ShowProgressWithMessage(z.ctl.UI(), MSortedWriter.ProgressPreparing)
	if err != nil {
		l.Warn("Unable to store row", esl.Error(err))
	}
}

func (z *Sorted) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) (err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}

	l := ctl.Log().With(esl.String("name", z.name))

	z.ctl = ctl
	z.dstPath = filepath.Join(ctl.Workspace().Report(), z.Name()+ro.ReportSuffix+"_column.json")
	z.dstFile, err = os.Create(z.dstPath)
	if err != nil {
		l.Debug("Unable to create a file", esl.Error(err))
		return err
	}
	z.sorter = es_sort.New(z.dstFile,
		es_sort.Logger(ctl.Log()),
		es_sort.TempCompress(true),
	)

	z.stream = rp_column_impl.NewStream(model, opts...)
	z.json = rp_column_impl.NewJson(z.stream.Header())

	newOpts := make([]rp_model.ReportOpt, 0)
	newOpts = append(newOpts, opts...)
	newOpts = append(newOpts, rp_model.ColumnModel(z.json))

	for _, w := range z.writers {
		if err := w.Open(ctl, model, newOpts...); err != nil {
			z.Close()
			return err
		}
	}
	z.isOpen = true
	return nil
}

func (z *Sorted) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	if z.ctl == nil {
		panic("the report is not yet opened")
	}
	l := z.ctl.Log()
	ui := z.ctl.UI()

	if !z.isOpen {
		l.Debug("The writer is not yet open")
		return
	}

	defer func() {
		l.Debug("Closing writers")
		for _, w := range z.writers {
			w.Close()
		}
	}()

	l.Debug("Writing sorted report")
	if clErr := z.sorter.Close(); clErr != nil {
		l.Debug("Unable to write sorted report", esl.Error(clErr))
		ui.Error(MSortedWriter.ErrorUnableComposeReport.With("Error", clErr))
		return
	}

	if z.dstFile == nil {
		l.Debug("Inconsistent state. Dst file is not found")
		ui.Error(MSortedWriter.ErrorUnableComposeReport.With("Error", "Failed to create the report file"))
		return
	}
	_ = z.dstFile.Close()

	df, err := os.Open(z.dstPath)
	if err != nil {
		l.Debug("Unable to open the dst file", esl.Error(err))
		ui.Error(MSortedWriter.ErrorUnableComposeReport.With("Error", err))
		return
	}
	dfs := bufio.NewScanner(df)
	var lastErr error
	for dfs.Scan() {
		line := dfs.Text()
		for _, w := range z.writers {
			w.Row([]byte(line))
		}
		app_ui.ShowProgressWithMessage(ui, MSortedWriter.ProgressSorting)
	}
	if err := dfs.Err(); err != nil {
		l.Debug("Error during read temporary report file", esl.Error(err))
		ui.Error(MSortedWriter.ErrorUnableComposeReport.With("Error", err))
	}
	if lastErr = dfs.Err(); lastErr != nil {
		ui.Error(MSortedWriter.ErrorUnableComposeReport.With("Error", lastErr))
	}

	_ = df.Close()
	if rmErr := os.Remove(z.dstPath); rmErr != nil {
		l.Debug("Unable to clean up column json", esl.Error(rmErr))
	}
}
