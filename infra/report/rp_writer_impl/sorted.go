package rp_writer_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_column_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sync"
)

type MsgSortedWriter struct {
	ProgressSorting   app_msg.Message
	ProgressPreparing app_msg.Message
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
	storage kv_storage.Storage
	stream  rp_column.Column
	bson    rp_column.Column
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

	err = z.storage.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutRaw(b, []byte{1})
	})
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
	z.storage, err = ctl.NewKvs("rp_writer_sorted-" + z.name + ro.ReportSuffix)
	if err != nil {
		l.Debug("Unable to create storage")
		return err
	}

	z.stream = rp_column_impl.NewStream(model, opts...)
	z.bson = rp_column_impl.NewBson(z.stream.Header())

	newOpts := make([]rp_model.ReportOpt, 0)
	newOpts = append(newOpts, opts...)
	newOpts = append(newOpts, rp_model.ColumnModel(z.bson))

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

	if !z.isOpen {
		l.Debug("The writer is not yet open")
		return
	}

	ui := z.ctl.UI()
	l.Debug("Writing sorted report")
	err := z.storage.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachRaw(func(key, value []byte) error {
			for _, w := range z.writers {
				w.Row(key)
			}
			app_ui.ShowProgressWithMessage(ui, MSortedWriter.ProgressSorting)
			return nil
		})
	})
	if err != nil {
		l.Debug("Unable to write sorted report", esl.Error(err))
	}

	for _, w := range z.writers {
		w.Close()
	}
	z.storage.Close()
}
