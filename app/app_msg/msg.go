package app_msg

import (
	"fmt"
	"github.com/watermint/toolbox/app/app_util"
	"go.uber.org/zap"
)

type UIMessage interface {
	WithArg(a ...interface{}) UIMessage
	WithTmplData(d interface{}) UIMessage
	Text() string
}

type Message struct {
	Message  string
	Args     []interface{}
	TmplData interface{}
	Logger   *zap.Logger
}

func (z *Message) WithTmplData(d interface{}) UIMessage {
	return &Message{
		Message:  z.Message,
		Args:     z.Args,
		TmplData: d,
		Logger:   z.Logger,
	}
}

func (z *Message) WithArg(a ...interface{}) UIMessage {
	return &Message{
		Message:  z.Message,
		Args:     a,
		TmplData: z.TmplData,
		Logger:   z.Logger,
	}
}

func (z *Message) Text() string {
	if z.TmplData != nil {
		t, err := app_util.CompileTemplate(z.Message, z.TmplData)
		if err != nil {
			z.Logger.Error(
				"Unable to compile template",
				zap.String("tmpl", z.Message),
				zap.Any("data", z.TmplData),
				zap.Error(err),
			)
			return z.Message
		}
		return t
	} else if z.Args != nil && len(z.Args) > 0 {
		return fmt.Sprintf(z.Message, z.Args...)
	} else {
		return z.Message
	}
}

func NewMessageMap(res map[string]string, log *zap.Logger) map[string]UIMessage {
	mm := make(map[string]UIMessage)
	for k, m := range res {
		mm[k] = &Message{
			Message: m,
			Logger:  log,
		}
	}
	return mm
}
