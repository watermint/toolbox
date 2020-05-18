package app_ui

import (
	"bytes"
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/strings/es_width"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"io"
	"strings"
)

func MakeConsoleDemo(mc app_msg_container.Container, f func(ui UI)) string {
	var buf bytes.Buffer
	lg := esl.Default()
	ui := NewProxy(
		&plainImpl{
			mc: mc,
			wr: &buf,
			dg: es_dialogue.DenyAll(),
		},
		lg,
	)
	f(ui)

	return buf.String()
}

type plainImpl struct {
	mc app_msg_container.Container
	wr io.Writer
	dg es_dialogue.Dialogue
}

func (z plainImpl) Header(m app_msg.Message) {
	h := z.mc.Compile(m)
	c := es_width.Width(h)

	z.Break()
	_, _ = fmt.Fprintln(z.wr, h)
	_, _ = fmt.Fprintln(z.wr, strings.Repeat("=", c))
	z.Break()
}

func (z plainImpl) SubHeader(m app_msg.Message) {
	h := z.mc.Compile(m)
	c := es_width.Width(h)

	z.Break()
	_, _ = fmt.Fprintln(z.wr, h)
	_, _ = fmt.Fprintln(z.wr, strings.Repeat("-", c))
	z.Break()
}

func (z plainImpl) Info(m app_msg.Message) {
	_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
}

func (z plainImpl) InfoTable(name string) Table {
	return newPlainTable(z, z.wr, z.mc, name)
}

func (z plainImpl) Error(m app_msg.Message) {
	_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
}

func (z plainImpl) Break() {
	_, _ = fmt.Fprintln(z.wr)
}

func (z plainImpl) AskProceed(m app_msg.Message) {
	z.dg.AskProceed(func() {
		_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
	})
}

func (z plainImpl) AskCont(m app_msg.Message) (cont bool) {
	p := func() {
		_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
	}
	return z.dg.AskCont(p, es_dialogue.YesNoCont)
}

func (z plainImpl) AskText(m app_msg.Message) (text string, cancel bool) {
	p := func() {
		_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
	}
	return z.dg.AskText(p, es_dialogue.NonEmptyText)
}

func (z plainImpl) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	p := func() {
		_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
	}
	return z.dg.AskSecure(p)
}

func (z plainImpl) Success(m app_msg.Message) {
	_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
}

func (z plainImpl) Failure(m app_msg.Message) {
	_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
}

func (z plainImpl) Progress(m app_msg.Message) {
	_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
}

func (z plainImpl) Quote(m app_msg.Message) {
	_, _ = fmt.Fprintln(z.wr, z.mc.Compile(m))
}

func (z plainImpl) Code(code string) {
	_, _ = fmt.Fprintln(z.wr, code)
}

func (z plainImpl) Link(artifact rp_artifact.Artifact) {
	z.Info(MConsole.PointArtifact.With("Path", artifact.Path()))
}

func (z plainImpl) IsConsole() bool {
	return true
}

func (z plainImpl) IsWeb() bool {
	return false
}

func (z plainImpl) WithContainerSyntax(mc app_msg_container.Container) Syntax {
	z.mc = mc
	return z
}

func (z plainImpl) Messages() app_msg_container.Container {
	return z.mc
}
