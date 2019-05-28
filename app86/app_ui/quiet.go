package app_ui

import (
	"github.com/watermint/toolbox/app86/app_msg"
	"go.uber.org/zap"
)

func NewQuiet() UI {
	return &Quiet{}
}

type Quiet struct {
	log *zap.Logger
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

func (z *Quiet) InfoTable(border bool) Table {
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

type QuietTable struct {
	log *zap.Logger
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
