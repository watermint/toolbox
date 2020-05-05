package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/strings/es_width"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"io"
	"strings"
)

func newMdTable(sy Syntax, wr io.Writer, mc app_msg_container.Container, name string) Table {
	return &mdTableImpl{
		sy:      sy,
		wr:      wr,
		mc:      mc,
		numRows: 0,
		name:    name,
		header:  []string{},
		rows:    make([][]string, 0),
	}
}

// Stateful:
type mdTableImpl struct {
	sy      Syntax
	wr      io.Writer
	mc      app_msg_container.Container
	numRows int
	name    string
	header  []string
	rows    [][]string
}

func (z *mdTableImpl) Header(h ...app_msg.Message) {
	hdr := make([]string, len(h))
	for i, m := range h {
		hdr[i] = z.mc.Compile(m)
	}
	z.HeaderRaw(hdr...)
}

func (z *mdTableImpl) HeaderRaw(h ...string) {
	z.header = h
}

func (z *mdTableImpl) Row(col ...app_msg.Message) {
	row := make([]string, len(col))
	for i, m := range col {
		row[i] = z.mc.Compile(m)
	}
	z.RowRaw(row...)
}

func (z *mdTableImpl) RowRaw(m ...string) {
	z.numRows++
	if z.numRows <= consoleNumRowsThreshold {
		z.rows = append(z.rows, m)
	}
	if z.numRows%consoleNumRowsThreshold == 0 {
		z.sy.Info(MConsole.Progress.
			With("Label", z.name).
			With("Progress", z.numRows))
	}
}

func (z *mdTableImpl) Flush() {
	l := es_log.Default()
	numCols := len(z.header)
	cols := make([]int, numCols)

	for i, c := range z.header {
		cols[i] = es_width.Width(c)
	}

	for _, row := range z.rows {
		rowNumCols := es_number.Min(len(row), numCols).Int()
		for i := 0; i < rowNumCols; i++ {
			cols[i] = es_number.Max(cols[i], es_width.Width(row[i])).Int()
		}
	}

	printCols := func(row []string) {
		_, _ = fmt.Fprintf(z.wr, "|")
		for i, c := range row {
			padding := 0
			if i < len(cols) {
				padding = es_number.Max(cols[i]-es_width.Width(c), 0).Int()
			} else {
				l.Debug("Number of columns exceeds header columns",
					es_log.Int("i", i),
					es_log.Strings("row", row))
				padding = 1
			}
			_, _ = fmt.Fprint(z.wr, " ")
			_, _ = fmt.Fprint(z.wr, c)
			_, _ = fmt.Fprint(z.wr, strings.Repeat(" ", padding))
			_, _ = fmt.Fprint(z.wr, " |")
		}
		_, _ = fmt.Fprintf(z.wr, "\n")
	}

	fmtBorder := "|"
	for _, c := range cols {
		fmtBorder += strings.Repeat("-", c+2) + "|"
	}

	printCols(z.header)
	_, _ = fmt.Fprintln(z.wr, fmtBorder)
	for _, row := range z.rows {
		printCols(row)
	}
	_, _ = fmt.Fprintln(z.wr, "")
}
