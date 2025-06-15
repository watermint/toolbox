package app_ui

import (
	"bytes"
	"fmt"
	"github.com/watermint/toolbox/essentials/io/es_line"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"io"
	"strings"
)

func MakeMarkdown(mc app_msg_container.Container, f func(ui UI)) string {
	var buf bytes.Buffer
	rlw := es_line.NewRemoveRedundantLinesWriter(&buf)
	lg := esl.Default()
	ui := NewProxy(
		&mdImpl{
			mc: mc,
			wr: rlw,
			dg: es_dialogue.DenyAll(),
		},
		lg,
	)
	f(ui)

	return buf.String()
}

func NewMarkdown(mc app_msg_container.Container, lg esl.Logger, wr io.Writer, dg es_dialogue.Dialogue) UI {
	return NewProxy(
		&mdImpl{
			mc: mc,
			wr: wr,
			dg: dg,
		},
		lg,
	)
}

type mdImpl struct {
	mc app_msg_container.Container
	wr io.Writer
	dg es_dialogue.Dialogue
}

func (z mdImpl) Messages() app_msg_container.Container {
	return z.mc
}

func (z mdImpl) Header(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "# %s\n\n", z.mc.Compile(m))
}

func (z mdImpl) SubHeader(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "\n## %s\n\n", z.mc.Compile(m))
}

func (z mdImpl) Info(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "%s\n", z.mc.Compile(m))
}

func (z mdImpl) InfoTable(name string) Table {
	return newMdTable(z, z.wr, z.mc, name)
}

func (z mdImpl) Error(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "**ERROR**: %s\n", z.mc.Compile(m))
}

func (z mdImpl) Quote(m app_msg.Message) {
	text := z.mc.Compile(m)
	for _, line := range strings.Split(text, "\n") {
		_, _ = fmt.Fprintf(z.wr, "> %s\n", line)
	}
}

func (z mdImpl) DefinitionList(definitions []Definition) {
	for _, def := range definitions {
		term := z.mc.Compile(def.Term)
		desc := z.mc.Compile(def.Description)
		
		// Use markdown-style definition list format
		_, _ = fmt.Fprintf(z.wr, "**%s**\n", term)
		_, _ = fmt.Fprintf(z.wr, ": %s\n\n", desc)
	}
}

func (z mdImpl) Break() {
	_, _ = fmt.Fprintf(z.wr, "\n")
}

func (z mdImpl) AskProceed(m app_msg.Message) {
	z.dg.AskProceed(func() {
		_, _ = fmt.Fprintf(z.wr, "_%s_\n", z.mc.Compile(m))
	})
}

func (z mdImpl) AskCont(m app_msg.Message) (cont bool) {
	p := func() {
		_, _ = fmt.Fprintf(z.wr, "_%s_\n", z.mc.Compile(m))
	}
	return z.dg.AskCont(p, es_dialogue.YesNoCont)
}

func (z mdImpl) AskText(m app_msg.Message) (text string, cancel bool) {
	p := func() {
		_, _ = fmt.Fprintf(z.wr, "_%s_\n", z.mc.Compile(m))
	}
	return z.dg.AskText(p, es_dialogue.NonEmptyText)
}

func (z mdImpl) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	p := func() {
		_, _ = fmt.Fprintf(z.wr, "_%s_\n", z.mc.Compile(m))
	}
	return z.dg.AskSecure(p)
}

func (z mdImpl) Success(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "**SUCCESS**: %s\n", z.mc.Compile(m))
}

func (z mdImpl) Failure(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "**FAILURE**: %s\n", z.mc.Compile(m))
}

func (z mdImpl) Progress(m app_msg.Message) {
	_, _ = fmt.Fprintf(z.wr, "_%s_\n", z.mc.Compile(m))
}

func (z mdImpl) Code(code string) {
	_, _ = fmt.Fprintf(z.wr, "```\n%s", code)
	if !strings.HasSuffix(code, "\n") {
		_, _ = fmt.Fprintf(z.wr, "\n")
	}
	_, _ = fmt.Fprintf(z.wr, "```\n")
}

func (z mdImpl) Link(artifact rp_artifact.Artifact) {
	_, _ = fmt.Fprintf(z.wr, "* [%s](%s)\n", artifact.Name(), artifact.Path())
}

func (z mdImpl) IsConsole() bool {
	return true
}

func (z mdImpl) IsWeb() bool {
	return false
}

func (z mdImpl) WithContainerSyntax(mc app_msg_container.Container) Syntax {
	z.mc = mc
	return z
}
