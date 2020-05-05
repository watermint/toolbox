package app_msg_container

import "github.com/watermint/toolbox/infra/ui/app_msg"

type Container interface {
	// Verify the key exists on the container.
	Exists(key string) bool

	// Compile message.
	Compile(m app_msg.Message) string

	// Pre compiled text for the key.
	Text(key string) string
}
