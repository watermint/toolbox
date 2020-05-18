package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/strings/es_width"
	"github.com/watermint/toolbox/essentials/terminal/es_color"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"io"
	"io/ioutil"
	"strings"
)

func NewDiscard(mc app_msg_container.Container, lg esl.Logger) UI {
	out := ioutil.Discard
	return NewConsole(
		mc,
		lg,
		out,
		es_dialogue.DenyAll(),
	)
}

func NewConsole(mc app_msg_container.Container, lg esl.Logger, wr io.Writer, dg es_dialogue.Dialogue) UI {
	return NewProxy(
		&conImpl{
			mc: mc,
			wr: wr,
			dg: dg,
		},
		lg,
	)
}

type MsgConsole struct {
	LargeReport       app_msg.Message
	OpenArtifactError app_msg.Message
	OpenArtifact      app_msg.Message
	PointArtifact     app_msg.Message
	Progress          app_msg.Message
}

var (
	MConsole = app_msg.Apply(&MsgConsole{}).(*MsgConsole)
)

type conImpl struct {
	mc app_msg_container.Container
	wr io.Writer
	dg es_dialogue.Dialogue
}

func (z conImpl) Messages() app_msg_container.Container {
	return z.mc
}

func (z conImpl) Header(m app_msg.Message) {
	h := z.mc.Compile(m)
	c := es_width.Width(h)

	z.Break()
	es_color.Boldfln(z.wr, h)
	es_color.Boldfln(z.wr, strings.Repeat("=", c))
	z.Break()
}

func (z conImpl) SubHeader(m app_msg.Message) {
	h := z.mc.Compile(m)
	c := es_width.Width(h)

	z.Break()
	es_color.Boldfln(z.wr, h)
	es_color.Boldfln(z.wr, strings.Repeat("-", c))
	z.Break()
}

func (z conImpl) Info(m app_msg.Message) {
	es_color.Colorfln(z.wr, es_color.ColorWhite, false, z.mc.Compile(m))
}

func (z conImpl) InfoTable(name string) Table {
	return newConTable(z, z.wr, z.mc, name)
}

func (z conImpl) Error(m app_msg.Message) {
	es_color.Colorfln(z.wr, es_color.ColorRed, false, z.mc.Compile(m))
}

func (z conImpl) Quote(m app_msg.Message) {
	es_color.Colorfln(z.wr, es_color.ColorGreen, false, z.mc.Compile(m))
}

func (z conImpl) Break() {
	_, _ = fmt.Fprintln(z.wr)
}

func (z conImpl) AskProceed(m app_msg.Message) {
	z.dg.AskProceed(func() {
		es_color.Colorfln(z.wr, es_color.ColorCyan, false, z.mc.Compile(m))
	})
}

func (z conImpl) AskCont(m app_msg.Message) (cont bool) {
	p := func() {
		es_color.Colorfln(z.wr, es_color.ColorCyan, false, z.mc.Compile(m))
	}
	return z.dg.AskCont(p, es_dialogue.YesNoCont)
}

func (z conImpl) AskText(m app_msg.Message) (text string, cancel bool) {
	p := func() {
		es_color.Colorfln(z.wr, es_color.ColorCyan, false, z.mc.Compile(m))
	}
	return z.dg.AskText(p, es_dialogue.NonEmptyText)
}

func (z conImpl) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	p := func() {
		es_color.Colorfln(z.wr, es_color.ColorCyan, false, z.mc.Compile(m))
	}
	return z.dg.AskSecure(p)
}

func (z conImpl) Success(m app_msg.Message) {
	es_color.Colorfln(z.wr, es_color.ColorGreen, false, z.mc.Compile(m))
}

func (z conImpl) Failure(m app_msg.Message) {
	es_color.Colorfln(z.wr, es_color.ColorRed, false, z.mc.Compile(m))
}

func (z conImpl) Progress(m app_msg.Message) {
	es_color.Colorfln(z.wr, es_color.ColorCyan, false, z.mc.Compile(m))
}

func (z conImpl) Code(code string) {
	es_color.Colorfln(z.wr, es_color.ColorBlue, false, code)
}

func (z conImpl) Link(artifact rp_artifact.Artifact) {
	z.Info(MConsole.PointArtifact.With("Path", artifact.Path()))
}

func (z conImpl) IsConsole() bool {
	return true
}

func (z conImpl) IsWeb() bool {
	return false
}

func (z conImpl) WithContainerSyntax(mc app_msg_container.Container) Syntax {
	z.mc = mc
	return z
}
