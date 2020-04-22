package app_msg_container_impl

import (
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg_impl"
	"go.uber.org/zap"
)

func NewMultilingual(las []lang.Lang, containers map[lang.Iso639One]app_msg_container.Container) app_msg_container.Container {
	return &Multilingual{
		LangPriority: las,
		Containers:   containers,
		qm:           qt_missingmsg_impl.NewMessageMemory(),
	}
}

type Multilingual struct {
	LangPriority []lang.Lang
	Containers   map[lang.Iso639One]app_msg_container.Container
	qm           qt_missingmsg.Message
}

func (z *Multilingual) Verify(key string) {
	for _, la := range z.LangPriority {
		if c, ok := z.Containers[la.Code()]; ok {
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
	for _, la := range z.LangPriority {
		if c, ok := z.Containers[la.Code()]; ok {
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
	for _, la := range z.LangPriority {
		if c, ok := z.Containers[la.Code()]; ok {
			if c.Exists(key) {
				return true
			}
		}
	}
	return false
}

func (z *Multilingual) Compile(m app_msg.Message) string {
	for _, la := range z.LangPriority {
		if c, ok := z.Containers[la.Code()]; ok {
			if c.Exists(m.Key()) {
				return c.Compile(m)
			}
		}
	}
	z.qm.NotFound(m.Key())
	app_root.Log().Warn("Unable to find message resource",
		zap.String("key", m.Key()),
	)

	alt := Alt{}
	return alt.Compile(m)
}
