package da_griddata

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"os"
	"path/filepath"
	"sync"
)

func NewPlainWriter(name, path string, formatter GridDataFormatter, writer PlainGridDataWriter) GridDataWriter {
	return &plainWriter{
		name:      name,
		path:      path,
		formatter: formatter,
		pw:        writer,
	}
}

type plainWriter struct {
	ctl       app_control.Control
	name      string
	formatter GridDataFormatter
	file      *os.File
	pw        PlainGridDataWriter
	lastErr   error
	path      string
	row       int
	mutex     sync.Mutex
}

func (z *plainWriter) Name() string {
	return z.name
}

func (z *plainWriter) Row(column []interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	l := z.ctl.Log().With(esl.String("path", z.path), esl.String("name", z.name))

	err := z.pw.WriteRow(l, z.file, z.formatter, z.row, column)
	if err != nil {
		if z.lastErr != err {
			l.Debug("Unable to write a row", esl.Error(err))
		}
		z.lastErr = err
	}

	// increment row index
	z.row++
}

func (z *plainWriter) Open(c app_control.Control) (err error) {
	z.ctl = c
	if z.path == "" {
		z.path = filepath.Join(c.Workspace().Report(), z.name+z.pw.FileSuffix())
	}
	l := c.Log().With(esl.String("path", z.path), esl.String("name", z.name))
	z.file, err = os.Create(z.path)
	if err != nil {
		l.Debug("Unable to create an output file", esl.Error(err))
		return err
	}
	return nil
}

func (z *plainWriter) Close() {
	l := z.ctl.Log().With(esl.String("name", z.name), esl.String("path", z.path))
	if z.file != nil {
		l.Debug("Closing file")
		_ = z.file.Close()
		z.file = nil
	} else {
		l.Debug("The file is not opened or already closed")
	}
}
