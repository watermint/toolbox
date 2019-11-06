package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/util/ut_math"
	"github.com/watermint/toolbox/infra/util/ut_string"
	"go.uber.org/zap"
	"io"
	"strings"
	"text/template"
)

func NewMarkdown(mc app_msg_container.Container, out io.Writer, ignoreMissing bool) UI {
	return &Markdown{
		mc:            mc,
		out:           out,
		ignoreMissing: ignoreMissing,
	}
}

type Markdown struct {
	mc            app_msg_container.Container
	out           io.Writer
	ignoreMissing bool
}

func (z *Markdown) print(tmpl, key string, p ...app_msg.P) {
	if z.ignoreMissing {
		if !z.mc.Exists(key) {
			return
		}
	}
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		app_root.Log().Warn("Template compile error", zap.String("template", tmpl), zap.Error(err))
		return
	}
	t.Execute(z.out, map[string]interface{}{
		"Message": z.mc.Compile(app_msg.M(key, p...)),
	})
}

func (z *Markdown) Header(key string, p ...app_msg.P) {
	z.print("# {{.Message}}\n\n", key, p...)
}

func (z *Markdown) Info(key string, p ...app_msg.P) {
	z.print("{{.Message}}\n", key, p...)
}

func (z *Markdown) InfoTable(name string) Table {
	return newMarkdownTable(z.mc, z.out, z.ignoreMissing)
}

func (z *Markdown) Error(key string, p ...app_msg.P) {
	z.print("ERROR: {{.Message}}\n", key, p...)
}

func (z *Markdown) Break() {
	fmt.Fprintln(z.out, "")
	fmt.Fprintln(z.out, "")
}

func (z *Markdown) Text(key string, p ...app_msg.P) string {
	return z.Text(key, p...)
}

func (z *Markdown) TextOrEmpty(key string, p ...app_msg.P) string {
	if z.mc.Exists(key) {
		return z.mc.Compile(app_msg.M(key, p...))
	} else {
		return ""
	}
}

func (z *Markdown) AskCont(key string, p ...app_msg.P) (cont bool, cancel bool) {
	return false, true
}

func (z *Markdown) AskText(key string, p ...app_msg.P) (text string, cancel bool) {
	return "", true
}

func (z *Markdown) AskSecure(key string, p ...app_msg.P) (secure string, cancel bool) {
	return "", true
}

func (z *Markdown) OpenArtifact(path string) {
}

func (z *Markdown) Success(key string, p ...app_msg.P) {
	z.print("SUCCESS: {{.Message}}\n", key, p...)
}

func (z *Markdown) Failure(key string, p ...app_msg.P) {
	z.print("FAILURE: {{.Message}}\n", key, p...)
}

func (z *Markdown) IsConsole() bool {
	return true
}

func (z *Markdown) IsWeb() bool {
	return false
}

func newMarkdownTable(mc app_msg_container.Container, out io.Writer, ignoreMissing bool) Table {
	return &markdownTable{
		mc:            mc,
		out:           out,
		ignoreMissing: ignoreMissing,
		header:        make([]string, 0),
		rows:          make([][]string, 0),
	}
}

type markdownTable struct {
	mc            app_msg_container.Container
	out           io.Writer
	ignoreMissing bool
	header        []string
	rows          [][]string
}

func (z *markdownTable) Header(h ...app_msg.Message) {
	z.header = make([]string, 0)
	for _, m := range h {
		if z.ignoreMissing {
			if !z.mc.Exists(m.Key()) {
				z.header = append(z.header, "")
				continue
			}
		}
		z.header = append(z.header, z.mc.Compile(m))
	}
}

func (z *markdownTable) HeaderRaw(h ...string) {
	z.header = make([]string, 0)
	z.header = append(z.header, h...)
}

func (z *markdownTable) Row(m ...app_msg.Message) {
	row := make([]string, 0)
	for _, n := range m {
		if z.ignoreMissing {
			if !z.mc.Exists(n.Key()) {
				row = append(row, "")
				continue
			}
		}
		row = append(row, z.mc.Compile(n))
	}
	z.RowRaw(row...)
}

func (z *markdownTable) RowRaw(m ...string) {
	z.rows = append(z.rows, m)
}

func (z *markdownTable) Flush() {
	numCols := len(z.header)
	cols := make([]int, numCols)

	for i, c := range z.header {
		cols[i] = ut_string.Width(c)
	}

	for _, row := range z.rows {
		rowNumCols := ut_math.MinInt(len(row), numCols)
		for i := 0; i < rowNumCols; i++ {
			cols[i] = ut_math.MaxInt(cols[i], ut_string.Width(row[i]))
		}
	}

	printCols := func(row []string) {
		fmt.Fprintf(z.out, "|")
		for i, c := range row {
			padding := ut_math.MaxInt(cols[i]-ut_string.Width(c), 0)
			fmt.Fprint(z.out, " ")
			fmt.Fprint(z.out, c)
			fmt.Fprint(z.out, strings.Repeat(" ", padding))
			fmt.Fprint(z.out, " |")
		}
		fmt.Fprintf(z.out, "\n")
	}

	fmtBorder := "|"
	for _, c := range cols {
		fmtBorder += strings.Repeat("-", c+2) + "|"
	}

	printCols(z.header)
	fmt.Fprintln(z.out, fmtBorder)
	for _, row := range z.rows {
		printCols(row)
	}
}
