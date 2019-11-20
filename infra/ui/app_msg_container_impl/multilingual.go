package app_msg_container_impl

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"github.com/watermint/toolbox/quality/infra/qt_control_impl"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

func NewMultilingual(langs []language.Tag, containers map[language.Tag]app_msg_container.Container) app_msg_container.Container {
	return &Multilingual{
		LangPriority: langs,
		Containers:   containers,
		qm:           qt_control_impl.NewMessageMemory(),
	}
}

type Multilingual struct {
	LangPriority []language.Tag
	Containers   map[language.Tag]app_msg_container.Container
	qm           qt_control.Message
}

func (z *Multilingual) Verify(key string) {
	for _, lang := range z.LangPriority {
		if c, ok := z.Containers[lang]; ok {
			if c.Exists(key) {
				return
			}
		}
	}
	z.qm.NotFound(key)
}

func (z *Multilingual) MissingKeys() []string {
	return z.qm.Missing()
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
