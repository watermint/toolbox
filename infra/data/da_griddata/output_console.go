package da_griddata

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"sync"
)

func NewConsoleWriter(formatter GridDataFormatter, pw PlainGridDataWriter) GridDataWriter {
	return &consoleWriter{
		formatter: formatter,
		pw:        pw,
	}
}

type consoleWriter struct {
	ctl       app_control.Control
	name      string
	formatter GridDataFormatter
	pw        PlainGridDataWriter
	row       int
	mutex     sync.Mutex
}

func (z *consoleWriter) Name() string {
	return z.name
}

func (z *consoleWriter) Row(column []interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	out := es_stdout.NewDefaultOut(z.ctl.Feature())

	_ = z.pw.WriteRow(z.ctl.Log(), out, z.formatter, z.row, column)
	// increment row index
	z.row++
}

func (z *consoleWriter) Open(c app_control.Control) error {
	z.ctl = c
	return nil
}

func (z *consoleWriter) Close() {
}
