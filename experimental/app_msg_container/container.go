package app_msg_container

import "github.com/watermint/toolbox/experimental/app_msg"

type Container interface {
	Exists(key string) bool
	Compile(m app_msg.Message) string
}
