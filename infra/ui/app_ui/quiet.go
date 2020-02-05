package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
)

func NewQuiet(container app_msg_container.Container) UI {
	q := container.(app_msg_container.Quality)
	return &Quiet{
		mc: container,
		mq: q,
	}
}

type Quiet struct {
	mc  app_msg_container.Container
	mq  app_msg_container.Quality
	log *zap.Logger
}

func (z *Quiet) AskCont(m app_msg.Message) (cont bool, cancel bool) {
	return false, true
}

func (z *Quiet) AskText(m app_msg.Message) (text string, cancel bool) {
	return "", true
}

func (z *Quiet) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	return "", true
}

func (z *Quiet) Header(m app_msg.Message) {
	z.mq.Verify(m.Key())
	z.log.Debug(m.Key(), zap.Any("params", m.Params()))
}

func (z *Quiet) Text(m app_msg.Message) string {
	z.mq.Verify(m.Key())
	z.log.Debug(m.Key(), zap.Any("params", m.Params()))
	return z.mc.Compile(m)
}

func (z *Quiet) TextOrEmpty(m app_msg.Message) string {
	if z.mc.Exists(m.Key()) {
		return z.mc.Compile(m)
	} else {
		return ""
	}
}

func (z *Quiet) Info(m app_msg.Message) {
	z.InfoK(m.Key(), m.Params()...)
}

func (z *Quiet) Error(m app_msg.Message) {
	z.ErrorK(m.Key(), m.Params()...)
}

func (z *Quiet) Success(m app_msg.Message) {
	z.mq.Verify(m.Key())
	z.log.Debug(m.Key(), zap.Any("params", m.Params()))
}

func (z *Quiet) Failure(m app_msg.Message) {
	z.mq.Verify(m.Key())
	z.log.Debug(m.Key(), zap.Any("params", m.Params()))
}

func (z *Quiet) SuccessK(key string, p ...app_msg.P) {
	z.Success(app_msg.M(key, p...))
}

func (z *Quiet) FailureK(key string, p ...app_msg.P) {
	z.Failure(app_msg.M(key, p...))
}

func (z *Quiet) IsConsole() bool {
	return true
}

func (z *Quiet) IsWeb() bool {
	return false
}

func (z *Quiet) OpenArtifact(path string, autoOpen bool) {
	z.log.Debug("Open artifact", zap.String("path", path))
}

func (z *Quiet) TextK(key string, p ...app_msg.P) string {
	z.mq.Verify(key)
	return z.mc.Compile(app_msg.M(key, p...))
}

func (z *Quiet) TextOrEmptyK(key string, p ...app_msg.P) string {
	if z.mc.Exists(key) {
		return z.mc.Compile(app_msg.M(key, p...))
	} else {
		return ""
	}
}
func (z *Quiet) SetLogger(log *zap.Logger) {
	z.log = log
}

func (z *Quiet) Break() {
	z.log.Debug("Break")
}

func (z *Quiet) HeaderK(key string, p ...app_msg.P) {
	z.mq.Verify(key)
	z.log.Debug(key, zap.Any("params", p))
}

func (z *Quiet) InfoTable(name string) Table {
	return &QuietTable{
		log: z.log,
		mq:  z.mq,
	}
}

func (z *Quiet) InfoK(key string, p ...app_msg.P) {
	z.mq.Verify(key)
	z.log.Debug(key, zap.Any("params", p))
}

func (z *Quiet) ErrorK(key string, p ...app_msg.P) {
	z.mq.Verify(key)
	z.log.Debug(key, zap.Any("params", p))
	z.log.Error(z.mc.Compile(app_msg.M(key, p...)))
}

// always cancel process
func (z *Quiet) AskContK(key string, p ...app_msg.P) (cont bool, cancel bool) {
	z.mq.Verify(key)
	z.log.Debug(key, zap.Any("params", p))
	return false, true
}

// always cancel
func (z *Quiet) AskTextK(key string, p ...app_msg.P) (text string, cancel bool) {
	z.mq.Verify(key)
	z.log.Debug(key, zap.Any("params", p))
	return "", true
}

// always cancel
func (z *Quiet) AskSecureK(key string, p ...app_msg.P) (secure string, cancel bool) {
	z.mq.Verify(key)
	z.log.Debug(key, zap.Any("params", p))
	return "", true
}

type QuietTable struct {
	log *zap.Logger
	mq  app_msg_container.Quality
}

func (z *QuietTable) HeaderRaw(h ...string) {
	z.log.Debug("header", zap.Any("h", h))
}

func (z *QuietTable) RowRaw(m ...string) {
	z.log.Debug("row", zap.Any("m", m))
}

func (z *QuietTable) Header(h ...app_msg.Message) {
	z.log.Debug("header", zap.Any("h", h))
	for _, m := range h {
		z.mq.Verify(m.Key())
	}
}

func (z *QuietTable) Row(m ...app_msg.Message) {
	z.log.Debug("row", zap.Any("m", m))
	for _, r := range m {
		z.mq.Verify(r.Key())
	}
}

func (z *QuietTable) Flush() {
	z.log.Debug("Flush")
}
