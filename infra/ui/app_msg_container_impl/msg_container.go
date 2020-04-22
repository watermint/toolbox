package app_msg_container_impl

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
)

func NewContainer(box *rice.Box) app_msg_container.Container {
	containers := make(map[lang.Iso639One]app_msg_container.Container)
	usrLang := lang.Detect(lang.Supported)
	priority := lang.Priority(usrLang)

	var lastErr error = nil

	for _, la := range priority {
		c, err := NewResource(la, box)
		if err != nil {
			lastErr = err
			continue
		}
		containers[la.Code()] = c
	}
	if len(containers) < 1 {
		app_root.Log().Error("No resources loaded", zap.Error(lastErr))
		panic("At least one message resource required")
	}

	return NewMultilingual(
		priority,
		containers,
	)
}
