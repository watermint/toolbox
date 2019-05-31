package app_msg_container_impl

import (
	"github.com/watermint/toolbox/atbx/app_msg"
	"github.com/watermint/toolbox/atbx/app_msg_container"
	"github.com/watermint/toolbox/atbx/app_root"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type Chain struct {
	LangPriority []language.Tag
	Containers   map[language.Tag]app_msg_container.Container
}

func (z *Chain) Exists(key string) bool {
	for _, lang := range z.LangPriority {
		if c, ok := z.Containers[lang]; ok {
			if c.Exists(key) {
				return true
			}
		}
	}
	return false
}

func (z *Chain) Compile(m app_msg.Message) string {
	for _, lang := range z.LangPriority {
		if c, ok := z.Containers[lang]; ok {
			if c.Exists(m.Key()) {
				return c.Compile(m)
			}
		}
	}
	app_root.Log().Warn("Unable to find message resource", zap.String("key", m.Key()))

	alt := Alt{}
	return alt.Compile(m)
}
