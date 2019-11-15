package app_msg_container

import "github.com/watermint/toolbox/infra/ui/app_msg"

type Container interface {
	Exists(key string) bool
	Compile(m app_msg.Message) string
	Text(key string) string
}

type Quality interface {
	Verify(key string)
	MissingKeys() []string
}
