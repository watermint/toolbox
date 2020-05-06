package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/terminal/es_color"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"io"
	"strings"
	"text/tabwriter"
)

func newConTable(sy Syntax, wr io.Writer, mc app_msg_container.Container, name string) Table {
	tw := new(tabwriter.Writer)
	tw.Init(wr, 0, 2, 2, ' ', 0)
	return &conTableImpl{
		sy:      sy,
		wr:      tw,
		mc:      mc,
		name:    name,
		limiter: NewRowLimiter(sy, name),
	}
}

type conTableImpl struct {
	sy      Syntax
	wr      *tabwriter.Writer
	mc      app_msg_container.Container
	name    string
	limiter RowLimiter
}

func (z conTableImpl) Header(h ...app_msg.Message) {
	headers := make([]string, 0)
	for _, hdr := range h {
		headers = append(headers, z.mc.Compile(hdr))
	}
	z.HeaderRaw(headers...)
}

func (z conTableImpl) HeaderRaw(h ...string) {
	es_color.Boldfln(z.wr, strings.Join(h, "\t"))
}

func (z conTableImpl) Row(m ...app_msg.Message) {
	ms := make([]string, 0)
	for _, msg := range m {
		ms = append(ms, z.mc.Compile(msg))
	}
	z.RowRaw(ms...)
}

func (z *conTableImpl) RowRaw(m ...string) {
	z.limiter.Row(func() {
		_, _ = fmt.Fprintln(z.wr, strings.Join(m, "\t"))
	})
}

func (z conTableImpl) Flush() {
	_ = z.wr.Flush()
	z.limiter.Flush()
}
