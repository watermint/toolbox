package rp_model_impl

import (
	"encoding/csv"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"sync"
)

func NewCsv(name string, row interface{}, ctl app_control.Control, opts ...rp_model.ReportOpt) (r rp_model.Report, err error) {
	l := ctl.Log()
	p := filepath.Join(ctl.Workspace().Report(), name+".csv")
	l.Debug("Create new csv report", zap.String("path", p))
	f, err := os.Create(p)
	if err != nil {
		l.Error("Unable to create file", zap.String("path", p), zap.Error(err))
		return nil, err
	}
	parser := NewColumn(row, opts...)
	r = &Csv{
		file:   f,
		w:      csv.NewWriter(f),
		ctl:    ctl,
		parser: parser,
	}
	return r, nil
}

type Csv struct {
	ctl    app_control.Control
	w      *csv.Writer
	file   *os.File
	mutex  sync.Mutex
	parser Column
	index  int
}

func (z *Csv) Success(input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSuccess.Key(), rp_model.MsgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Csv) Failure(reason app_msg.Message, input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgFailure.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Csv) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSkip.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Csv) Row(row interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.index == 0 {
		z.w.Write(z.parser.Header())
	}
	z.w.Write(z.parser.ValuesAsString(row))
	z.index++
}

func (z *Csv) flush() {
	z.w.Flush()
	z.file.Sync()
}

func (z *Csv) Close() {
	if z.file != nil {
		z.flush()
		z.file.Close()
		z.file = nil
	}
}
