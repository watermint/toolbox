package app_control_impl

import (
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"github.com/watermint/toolbox/infra/recipe/rc_worker_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func New(wb app_workspace.Bundle, ui app_ui.UI, feature app_feature.Feature) app_control.Control {
	return &ctlImpl{
		wb:      wb,
		ui:      ui,
		feature: feature,
	}
}

type ctlImpl struct {
	feature app_feature.Feature
	ui      app_ui.UI
	wb      app_workspace.Bundle
}

func (z ctlImpl) WithLang(targetLang string) app_control.Control {
	l := z.Log().With(es_log.String("targetLang", targetLang))
	usrLang := lang.Select(targetLang, lang.Supported)
	priority := lang.Priority(usrLang)
	containers := make(map[lang.Iso639One]app_msg_container.Container)

	for _, la := range priority {
		mc, err := app_msg_container_impl.NewSingle(la)
		if err != nil {
			l.Debug("Unable to load resource for language",
				es_log.String("la", la.String()),
				es_log.Error(err))
			return z
		}
		containers[la.Code()] = mc
	}

	mc := app_msg_container_impl.NewMultilingual(priority, containers)
	z.ui = z.ui.WithContainer(mc)
	return z
}

func (z ctlImpl) WithFeature(feature app_feature.Feature) app_control.Control {
	z.feature = feature
	return z
}

func (z ctlImpl) WithUI(ui app_ui.UI) app_control.Control {
	z.ui = ui
	return z
}

func (z ctlImpl) WithBundle(wb app_workspace.Bundle) app_control.Control {
	z.wb = wb
	return z
}

func (z ctlImpl) Feature() app_feature.Feature {
	return z.feature
}

func (z ctlImpl) Messages() app_msg_container.Container {
	return z.ui.Messages()
}

func (z ctlImpl) NewQueue() rc_worker.Queue {
	return rc_worker_impl.NewQueue(z, z.feature.Concurrency())
}

func (z ctlImpl) Workspace() app_workspace.Workspace {
	return z.wb.Workspace()
}

func (z ctlImpl) WorkBundle() app_workspace.Bundle {
	return z.wb
}

func (z ctlImpl) UI() app_ui.UI {
	return z.ui
}

func (z ctlImpl) Log() es_log.Logger {
	return z.wb.Logger().Logger()
}

func (z ctlImpl) Capture() es_log.Logger {
	return z.wb.Capture().Logger()
}
