package app_msg_container_impl

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"sync"
)

type Bundle map[string]*rice.Box

type bundleResource struct {
	// mutex lock for updating message resources
	mutex sync.Mutex

	// file name suffix
	langSuffix string

	// loaded resources
	messages map[string]string

	// true when the key already tried lazy load, but not found.
	blackList map[string]bool

	// relative path from toolbox package root.
	bundlePath string

	// box to the resources.
	box *rice.Box
}

func (z *bundleResource) Exists(key string) bool {
	panic("implement me")
}

func (z *bundleResource) Compile(m app_msg.Message) string {
	panic("implement me")
}

func (z *bundleResource) WithPrefix(prefix string) app_msg_container.Container {
	panic("implement me")
}
