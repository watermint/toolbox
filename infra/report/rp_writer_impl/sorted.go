package rp_writer_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_column_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
)

type MsgSortedWriter struct {
	ProgressSorting app_msg.Message
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
}

func (z *Sorted) Name() string {
	return z.name
}

func (z *Sorted) Row(r interface{}) {
	l := z.ctl.Log().With(zap.String("name", z.name))
	vals := z.stream.Values(r)

	b, err := json.Marshal(vals)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return
	}

	err = z.storage.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutRaw(b, []byte{1})
	})
	if err != nil {
		l.Debug("Unable to store row", zap.Error(err))
	}
}

func (z *Sorted) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
	z.ctl = ctl
	z.storage = kv_storage_impl.New("rp_writer_sorted-" + z.name)

	l := ctl.Log().With(zap.String("name", z.name))

	if err := z.storage.Open(ctl); err != nil {
		l.Debug("Unable to create storage")
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
	return nil
}

func (z *Sorted) Close() {
	l := z.ctl.Log()
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
		l.Debug("Unable to write sorted report", zap.Error(err))
	}

	for _, w := range z.writers {
		w.Close()
	}
	z.storage.Close()
}
