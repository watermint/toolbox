package app_msg_container_impl

import (
	"strings"

	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_msgusage"
)

func NewMultilingual(las []es_lang.Lang, containers map[es_lang.Iso639One]app_msg_container.Container) app_msg_container.Container {
	return &mlContainer{
		priority:   las,
		containers: containers,
	}
}

type mlContainer struct {
	priority   []es_lang.Lang
	containers map[es_lang.Iso639One]app_msg_container.Container
}

func (z mlContainer) Lang() es_lang.Lang {
	return z.priority[0]
}

func (z mlContainer) Text(key string) string {
	l := esl.Default()
	for _, la := range z.priority {
		if c, ok := z.containers[la.Code()]; ok {
			if c.ExistsKey(key) {
				return c.Text(key)
			}
		}
	}
	qt_msgusage.Record().NotFound(key)
	l.Warn("Unable to find message resource", esl.String("key", key))
	return AltText(key)
}

func (z *mlContainer) Exists(msg app_msg.Message) bool {
	switch m := msg.(type) {
	case app_msg.MessageComplex:
		for _, mm := range m.Messages() {
			if !z.ExistsKey(mm.Key()) {
				return false
			}
		}
		return true

	default:
		for _, la := range z.priority {
			if c, ok := z.containers[la.Code()]; ok {
				if c.Exists(msg) {
					return true
				}
			}
		}
		return false
	}
}

func (z *mlContainer) ExistsKey(key string) bool {
	for _, la := range z.priority {
		if c, ok := z.containers[la.Code()]; ok {
			if c.ExistsKey(key) {
				return true
			}
		}
	}
	return false
}

func (z *mlContainer) compileComplex(messages []app_msg.Message) string {
	compiled := make([]string, 0)
	for _, msg := range messages {
		compiled = append(compiled, z.Compile(msg))
	}
	return strings.Join(compiled, " ")
}

func (z *mlContainer) Compile(m app_msg.Message) string {
	l := esl.Default()
	key := m.Key()
	switch m0 := m.(type) {
	case app_msg.MessageComplex:
		qt_msgusage.Record().Touch(key)
		return z.compileComplex(m0.Messages())

	case app_msg.MessageOptional:
		for _, la := range z.priority {
			if c, ok := z.containers[la.Code()]; ok {
				if c.Exists(m) {
					qt_msgusage.Record().Touch(key)
					return c.Compile(m)
				}
			}
		}
		if m0.Optional() {
			qt_msgusage.Record().Touch(key)
			return ""
		} else {
			qt_msgusage.Record().NotFound(key)
			l.Warn("Unable to find message resource", esl.String("key", key))
			return AltCompile(m)
		}

	default:
		for _, la := range z.priority {
			if c, ok := z.containers[la.Code()]; ok {
				if c.Exists(m) {
					qt_msgusage.Record().Touch(key)
					return c.Compile(m)
				}
			}
		}
		qt_msgusage.Record().NotFound(key)
		l.Warn("Unable to find message resource", esl.String("key", key))
		return AltCompile(m)
	}
}
