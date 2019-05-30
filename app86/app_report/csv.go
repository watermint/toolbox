package app_report

import (
	"encoding/csv"
	"github.com/watermint/toolbox/app86/app_control"
	"os"
	"path/filepath"
)

func NewCsv(name string, ctl app_control.Control) (r Report, err error) {
	p, err := ctl.Workspace().Descendant(reportPath)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(filepath.Join(p, name+".csv"))
	if err != nil {
		return nil, err
	}
	r = &Csv{
		File:   f,
		Writer: csv.NewWriter(f),
		Ctl:    ctl,
	}
	return r, nil
}

type Csv struct {
	Ctl    app_control.Control
	Writer *csv.Writer
	File   *os.File
	Parser Row
}

func (z *Csv) Row(row interface{}) {
	if z.Parser == nil {
		z.Parser = NewRow(row, z.Ctl)
		z.Writer.Write(z.Parser.Header())
	}
	z.Writer.Write(z.Parser.ValuesAsString(row))
}

func (z *Csv) Transaction(state State, row interface{}, result interface{}) {
	z.Row(Transaction{State: state(), Input: row, Result: result})
}

func (z *Csv) Flush() {
	z.Writer.Flush()
	z.File.Sync()
}

func (z *Csv) Close() {
	if z.File != nil {
		z.Flush()
		z.File.Close()
		z.File = nil
	}
}
