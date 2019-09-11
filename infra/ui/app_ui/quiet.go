package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
)

func NewQuiet(container app_msg_container.Container) UI {
	return &Quiet{container: container}
}

type Quiet struct {
	container app_msg_container.Container
	log       *zap.Logger
}

func (z *Quiet) Success(key string, p ...app_msg.Param) {
	z.log.Debug(key, zap.Any("params", p))
}

func (z *Quiet) Failure(key string, p ...app_msg.Param) {
	z.log.Debug(key, zap.Any("params", p))
}

func (z *Quiet) IsConsole() bool {
	return false
}

func (z *Quiet) IsWeb() bool {
	return false
}

func (z *Quiet) OpenArtifact(path string) {
	z.log.Debug("Open artifact", zap.String("path", path))
}

func (z *Quiet) Text(key string, p ...app_msg.Param) string {
	return z.container.Compile(app_msg.M(key, p...))
}

func (z *Quiet) SetLogger(log *zap.Logger) {
	z.log = log
}

func (z *Quiet) Break() {
	z.log.Debug("Break")
}

func (z *Quiet) Header(key string, p ...app_msg.Param) {
	z.log.Debug(key, zap.Any("params", p))
}

func (z *Quiet) InfoTable(name string) Table {
	return &QuietTable{
		log: z.log,
	}
}

func (z *Quiet) Info(key string, p ...app_msg.Param) {
	z.log.Debug(key, zap.Any("params", p))
}

func (z *Quiet) Error(key string, p ...app_msg.Param) {
	z.log.Debug(key, zap.Any("params", p))
}

// always cancel process
func (z *Quiet) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	z.log.Debug(key, zap.Any("params", p))
	return false, true
}

// always cancel
func (z *Quiet) AskText(key string, p ...app_msg.Param) (text string, cancel bool) {
	z.log.Debug(key, zap.Any("params", p))
	return "", true
}

// always cancel
func (z *Quiet) AskSecure(key string, p ...app_msg.Param) (secure string, cancel bool) {
	z.log.Debug(key, zap.Any("params", p))
	return "", true
}

type QuietTable struct {
	log *zap.Logger
}

func (z *QuietTable) HeaderRaw(h ...string) {
	z.log.Debug("header", zap.Any("h", h))
}

func (z *QuietTable) RowRaw(m ...string) {
	z.log.Debug("row", zap.Any("m", m))
}

func (z *QuietTable) Header(h ...app_msg.Message) {
	z.log.Debug("header", zap.Any("h", h))
}

func (z *QuietTable) Row(m ...app_msg.Message) {
	z.log.Debug("row", zap.Any("m", m))
}

func (z *QuietTable) Flush() {
	z.log.Debug("Flush")
}
