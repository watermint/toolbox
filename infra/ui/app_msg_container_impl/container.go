package app_msg_container_impl

import (
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
)

func NewContainer() app_msg_container.Container {
	containers := make(map[lang.Iso639One]app_msg_container.Container)
	usrLang := lang.Detect(lang.Supported)
	priority := lang.Priority(usrLang)

	var lastErr error = nil

	for _, la := range priority {
		c, err := NewSingle(la)
		if err != nil {
			lastErr = err
			continue
		}
		containers[la.Code()] = c
	}
	if len(containers) < 1 {
		es_log.Default().Error("No resources loaded", es_log.Error(lastErr))
		panic("At least one message resource required")
	}

	return NewMultilingual(
		priority,
		containers,
	)
}
