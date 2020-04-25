package context

import (
	"go.uber.org/zap"
)

type Context interface {
	ClientHash() string
	Log() *zap.Logger
	Capture() *zap.Logger
}
