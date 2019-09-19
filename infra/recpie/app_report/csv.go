package app_report

import (
	"encoding/csv"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
	"sync"
)

func NewCsv(name string, row interface{}, ctl app_control.Control) (r Report, err error) {
	p, err := ctl.Workspace().Descendant(reportPath)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(filepath.Join(p, name+".csv"))
	if err != nil {
		return nil, err
	}
	parser := NewColumn(row, ctl)
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
	z.Row(TransactionRow{
		Status: ui.Text(msgSuccess.Key(), msgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Csv) Failure(reason app_msg.Message, input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(TransactionRow{
		Status: ui.Text(msgFailure.Key(), msgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Csv) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(TransactionRow{
		Status: ui.Text(msgSkip.Key(), msgFailure.Params()...),
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

func (z *Csv) Flush() {
	z.w.Flush()
	z.file.Sync()
}

func (z *Csv) Close() {
	if z.file != nil {
		z.Flush()
		z.file.Close()
		z.file = nil
	}
}
