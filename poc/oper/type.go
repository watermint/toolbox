package oper

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/poc/oper/oper_msg"
	"github.com/watermint/toolbox/poc/oper/oper_ui"
	"go.uber.org/zap"
)

type Operation interface {
}

type Executable interface {
	Exec()
}

type Group interface {
	Operations() []Operation
}

type Resource struct {
	Title    string            `json:"title,omitempty"`
	Desc     string            `json:"desc,omitempty"`
	Options  map[string]string `json:"options,omitempty"`
	Messages map[string]string `json:"messages,omitempty"`
}

type Context interface {
	Log() *zap.Logger
	Box() *rice.Box
	UI() oper_ui.UI
	Message(key string) oper_msg.UIMessage
	WithMessages(messages map[string]oper_msg.UIMessage) Context
}

func NewContext(log *zap.Logger, box *rice.Box, ui oper_ui.UI) Context {
	return &contextImpl{
		logger: log,
		box:    box,
		ui:     ui,
	}
}

type contextImpl struct {
	logger   *zap.Logger
	box      *rice.Box
	ui       oper_ui.UI
	messages map[string]oper_msg.UIMessage
}

func (z *contextImpl) WithMessages(messages map[string]oper_msg.UIMessage) Context {
	return &contextImpl{
		logger:   z.logger,
		box:      z.box,
		ui:       z.ui,
		messages: messages,
	}
}

func (z *contextImpl) Log() *zap.Logger {
	return z.logger
}

func (z *contextImpl) Box() *rice.Box {
	return z.box
}

func (z *contextImpl) UI() oper_ui.UI {
	return z.ui
}

func (z *contextImpl) Message(key string) oper_msg.UIMessage {
	if m, ok := z.messages[key]; ok {
		return m
	} else {
		z.Log().Error("Message not found", zap.String("key", key))
		return &oper_msg.Message{
			Message: key,
		}
	}
}

type OperationBase struct {
	Ctx Context
}

func (z *OperationBase) Log() *zap.Logger {
	return z.Ctx.Log()
}

func (z *OperationBase) Message(key string) oper_msg.UIMessage {
	return z.Ctx.Message(key)
}

func (z *OperationBase) Tell(msg oper_msg.UIMessage) {
	z.Ctx.UI().Tell(msg)
}

func (z *OperationBase) TellError(msg oper_msg.UIMessage) {
	z.Ctx.UI().TellError(msg)
}

func (z *OperationBase) TellDone(msg oper_msg.UIMessage) {
	z.Ctx.UI().TellDone(msg)
}

func (z *OperationBase) TellSuccess(msg oper_msg.UIMessage) {
	z.Ctx.UI().TellSuccess(msg)
}

func (z *OperationBase) TellFailure(msg oper_msg.UIMessage) {
	z.Ctx.UI().TellFailure(msg)
}

func (z *OperationBase) TellProgress(msg oper_msg.UIMessage) {
	z.Ctx.UI().TellProgress(msg)
}

func (z *OperationBase) AskRetry(msg oper_msg.UIMessage) bool {
	return z.Ctx.UI().AskRetry(msg)
}

func (z *OperationBase) AskWarn(msg oper_msg.UIMessage) bool {
	return z.Ctx.UI().AskWarn(msg)
}

func (z *OperationBase) AskText(msg oper_msg.UIMessage) string {
	return z.Ctx.UI().AskText(msg)
}
