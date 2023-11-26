package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
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
		limiter: NewRowLimiter(sy, name),
		name:    name,
		header:  []string{},
		rows:    make([][]string, 0),
	}
}

type mdTableImpl struct {
	sy      Syntax
	wr      io.Writer
	mc      app_msg_container.Container
	limiter RowLimiter
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
	z.limiter.Row(func() {
		z.rows = append(z.rows, m)
	})
}

func (z *mdTableImpl) Flush() {
	l := esl.Default()
	numCols := len(z.header)
	cols := make([]int, numCols)

	for i, c := range z.header {
		cols[i] = es_width.Width(c)
	}

	for _, row := range z.rows {
		rowNumCols := min(len(row), numCols)
		for i := 0; i < rowNumCols; i++ {
			cols[i] = max(cols[i], es_width.Width(row[i]))
		}
	}

	printCols := func(row []string) {
		_, _ = fmt.Fprintf(z.wr, "|")
		for i, c := range row {
			padding := 0
			if i < len(cols) {
				padding = max(cols[i]-es_width.Width(c), 0)
			} else {
				l.Debug("Number of columns exceeds header columns",
					esl.Int("i", i),
					esl.Strings("row", row))
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

	_, _ = fmt.Fprintln(z.wr)
	printCols(z.header)
	_, _ = fmt.Fprintln(z.wr, fmtBorder)
	for _, row := range z.rows {
		printCols(row)
	}
	_, _ = fmt.Fprintln(z.wr, "")
	_, _ = fmt.Fprintln(z.wr)

	z.limiter.Flush()
}
