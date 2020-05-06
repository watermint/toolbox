package app_ui

import (
	"github.com/watermint/toolbox/essentials/concurrency/es_mutex"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
)

func NewProxy(sy Syntax, lg es_log.Logger) UI {
	id := newId()
	return &proxyImpl{
		id: id,
		lg: lg.With(es_log.String("id", id)),
		sy: sy,
		mx: es_mutex.New(),
	}
}

type proxyImpl struct {
	id string
	lg es_log.Logger
	sy Syntax
	mx es_mutex.Mutex
}

func (z proxyImpl) Messages() app_msg_container.Container {
	return z.sy.Messages()
}

func (z proxyImpl) WithTable(name string, f func(t Table)) {
	t := z.sy.InfoTable(name)
	f(t)
	t.Flush()
}

func (z proxyImpl) verifyKey(m app_msg.Message) {
	k := m.Key()
	if !z.sy.Messages().Exists(k) {
		z.lg.Debug("Message key not found", es_log.String("key", k))
		qt_missingmsg.Record().NotFound(k)
	}
}

func (z proxyImpl) withMsg(action string, m app_msg.Message, f func()) {
	z.verifyKey(m)
	z.lg.Debug(
		z.TextOrEmpty(m),
		es_log.String("action", action),
		es_log.String("key", m.Key()),
		es_log.Any("params", m.Params()))
	z.mx.Do(f)
}

func (z proxyImpl) Header(m app_msg.Message) {
	z.withMsg("header", m, func() {
		z.sy.Header(m)
	})
}

func (z proxyImpl) SubHeader(m app_msg.Message) {
	z.withMsg("subHeader", m, func() {
		z.sy.SubHeader(m)
	})
}

func (z proxyImpl) Info(m app_msg.Message) {
	z.withMsg("info", m, func() {
		z.sy.Info(m)
	})
}

func (z proxyImpl) InfoTable(name string) Table {
	return proxyTableImpl{
		lg: z.lg.With(es_log.String("name", name)),
		mc: z.sy.Messages(),
		mx: es_mutex.New(),
		it: z.sy.InfoTable(name),
	}
}

func (z proxyImpl) Error(m app_msg.Message) {
	z.withMsg("error", m, func() {
		z.sy.Error(m)
	})
}

func (z proxyImpl) Break() {
	z.mx.Do(func() {
		z.sy.Break()
	})
}

func (z proxyImpl) Exists(m app_msg.Message) bool {
	return z.sy.Messages().Exists(m.Key())
}

func (z proxyImpl) Text(m app_msg.Message) string {
	return z.sy.Messages().Compile(m)
}

func (z proxyImpl) TextOrEmpty(m app_msg.Message) string {
	if z.sy.Messages().Exists(m.Key()) {
		return z.sy.Messages().Compile(m)
	} else {
		return ""
	}
}

func (z proxyImpl) AskProceed(m app_msg.Message) {
	z.withMsg("askProceed", m, func() {
		z.sy.AskProceed(m)
	})
}

func (z proxyImpl) AskCont(m app_msg.Message) (cont bool) {
	z.withMsg("askCont", m, func() {
		cont = z.sy.AskCont(m)
	})
	return
}

func (z proxyImpl) AskText(m app_msg.Message) (text string, cancel bool) {
	z.withMsg("askText", m, func() {
		text, cancel = z.sy.AskText(m)
	})
	return
}

func (z proxyImpl) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	z.withMsg("askSecure", m, func() {
		secure, cancel = z.sy.AskSecure(m)
	})
	return
}

func (z proxyImpl) Success(m app_msg.Message) {
	z.withMsg("success", m, func() {
		z.sy.Success(m)
	})
}

func (z proxyImpl) Failure(m app_msg.Message) {
	z.withMsg("failure", m, func() {
		z.sy.Failure(m)
	})
}

func (z proxyImpl) Progress(m app_msg.Message) {
	z.withMsg("progress", m, func() {
		z.sy.Progress(m)
	})
}

func (z proxyImpl) Code(code string) {
	z.lg.Debug("code", es_log.String("code", code))
	z.mx.Do(func() {
		z.sy.Code(code)
	})
}

func (z proxyImpl) Link(artifact rp_artifact.Artifact) {
	z.lg.Debug("artifact", es_log.Any("artifact", artifact))
	z.mx.Do(func() {
		z.sy.Link(artifact)
	})
}

func (z proxyImpl) IsConsole() bool {
	return z.sy.IsConsole()
}

func (z proxyImpl) IsWeb() bool {
	return z.sy.IsWeb()
}

func (z proxyImpl) Id() string {
	return z.id
}

func (z proxyImpl) WithContainerSyntax(mc app_msg_container.Container) Syntax {
	return z.sy.WithContainerSyntax(mc)
}

func (z proxyImpl) WithContainer(mc app_msg_container.Container) UI {
	z.id = newId()
	z.sy = z.sy.WithContainerSyntax(mc)
	return z
}
