package pipe_core

import "go.uber.org/zap"

// Mutable?
type Service interface {
	Exec(req Request, s Session) (res Response, err error)
}

// Mutable
type Request interface {
}

// Mutable
type Response interface {
}

// Mutable?
type Session interface {
	Log() *zap.Logger
	Message(key string) UIMessage
	Set(key string, value interface{})
	Get(key string) (v interface{}, exists bool)
}

// Immutable
type UIMessage interface {
	WithParam(p interface{}) UIMessage
	Tell()
}
