package app_msg_container

import (
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Container interface {
	// Verify the key exists on the container.
	// This does not compute message type such as app_msg.MessageComplex.
	ExistsKey(key string) bool

	// Verify the message and keys exists on the container.
	Exists(msg app_msg.Message) bool

	// Compile message.
	Compile(m app_msg.Message) string

	// Pre compiled text for the key.
	Text(key string) string

	// Current language
	Lang() es_lang.Lang
}
