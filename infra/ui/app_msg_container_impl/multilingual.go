package app_msg_container_impl

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type Multilingual struct {
	LangPriority []language.Tag
	Containers   map[language.Tag]app_msg_container.Container
}

func (z *Multilingual) Text(key string) string {
	for _, lang := range z.LangPriority {
		if c, ok := z.Containers[lang]; ok {
			if c.Exists(key) {
				return c.Text(key)
			}
		}
	}
	app_root.Log().Warn("Unable to find message resource",
		zap.String("key", key),
	)

	alt := Alt{}
	return alt.Text(key)
}

func (z *Multilingual) Exists(key string) bool {
	for _, lang := range z.LangPriority {
		if c, ok := z.Containers[lang]; ok {
			if c.Exists(key) {
				return true
			}
		}
	}
	return false
}

func (z *Multilingual) Compile(m app_msg.Message) string {
	for _, lang := range z.LangPriority {
		if c, ok := z.Containers[lang]; ok {
			if c.Exists(m.Key()) {
				return c.Compile(m)
			}
		}
	}
	app_root.Log().Warn("Unable to find message resource",
		zap.String("key", m.Key()),
	)

	alt := Alt{}
	return alt.Compile(m)
}
