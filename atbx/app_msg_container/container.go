package app_msg_container

import "github.com/watermint/toolbox/atbx/app_msg"

type Container interface {
	Exists(key string) bool
	Compile(m app_msg.Message) string
}
