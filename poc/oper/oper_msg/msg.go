package oper_msg

import "fmt"

type UIMessage interface {
	WithArg(a ...interface{}) UIMessage
	Text() string
}

type Message struct {
	Message string
	Args    []interface{}
}

func (z *Message) WithArg(a ...interface{}) UIMessage {
	return &Message{
		Message: z.Message,
		Args:    a,
	}
}

func (z *Message) Text() string {
	if z.Args == nil {
		return z.Message
	} else {
		return fmt.Sprintf(z.Message, z.Args...)
	}
}

func NewMessageMap(res map[string]string) map[string]UIMessage {
	mm := make(map[string]UIMessage)
	for k, m := range res {
		mm[k] = &Message{
			Message: m,
		}
	}
	return mm
}
