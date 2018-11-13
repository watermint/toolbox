package oper_i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type UIMessage interface {
	Message(lang language.Tag) string
}

type ResourceMessage struct {
	Key  string
	Args []interface{}
}

func (z *ResourceMessage) Message(lang language.Tag) string {
	p := message.NewPrinter(lang)
	return p.Sprintf(z.Key, z.Args...)
}

type Messages struct {
	KeyText map[string]string
}

func (z *Messages) Msg(key string, arg ...interface{}) UIMessage {
	return &ResourceMessage{
		Key:  key,
		Args: arg,
	}
}
