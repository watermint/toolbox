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

type Context struct {
	Logger   *zap.Logger
	Box      *rice.Box
	UI       oper_ui.UI
	Messages map[string]oper_msg.UIMessage
}

type OperationBase struct {
	Ctx *Context
}

func (z *OperationBase) Log() *zap.Logger {
	return z.Ctx.Logger
}

func (z *OperationBase) Message(key string) oper_msg.UIMessage {
	if m, ok := z.Ctx.Messages[key]; ok {
		return m
	} else {
		z.Ctx.Logger.Error("Message not found", zap.String("key", key))
		return &oper_msg.Message{
			Message: key,
		}
	}
}

func (z *OperationBase) Tell(msg oper_msg.UIMessage) {
	z.Ctx.UI.Tell(msg)
}

func (z *OperationBase) TellError(msg oper_msg.UIMessage) {
	z.Ctx.UI.TellError(msg)
}

func (z *OperationBase) TellDone(msg oper_msg.UIMessage) {
	z.Ctx.UI.TellDone(msg)
}

func (z *OperationBase) TellSuccess(msg oper_msg.UIMessage) {
	z.Ctx.UI.TellSuccess(msg)
}

func (z *OperationBase) TellFailure(msg oper_msg.UIMessage) {
	z.Ctx.UI.TellFailure(msg)
}

func (z *OperationBase) TellProgress(msg oper_msg.UIMessage) {
	z.Ctx.UI.TellProgress(msg)
}

func (z *OperationBase) AskRetry(msg oper_msg.UIMessage) bool {
	return z.Ctx.UI.AskRetry(msg)
}

func (z *OperationBase) AskWarn(msg oper_msg.UIMessage) bool {
	return z.Ctx.UI.AskWarn(msg)
}

func (z *OperationBase) AskText(msg oper_msg.UIMessage) string {
	return z.Ctx.UI.AskText(msg)
}
