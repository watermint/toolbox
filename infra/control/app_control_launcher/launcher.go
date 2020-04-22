package app_control_launcher

import (
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
)

type ControlLauncher interface {
	Catalogue() rc_catalogue.Catalogue
	NewControl(user app_workspace.MultiUser) (ctl app_control.Control, err error)
}

// Fork control: create workspace with name under existing control
type ControlFork interface {
	Fork(name string, opts ...app_control.UpOpt) (ctl app_control.Control, err error)
}

type WithMessageContainer interface {
	With(mc app_msg_container.Container) app_control.Control
}

func ControlWithLang(targetLang string, c app_control.Control) (app_control.Control, bool) {
	wc, ok := c.(WithMessageContainer)
	if !ok {
		return nil, false
	}

	usrLang := lang.Select(targetLang, lang.Supported)
	priority := lang.Priority(usrLang)
	containers := make(map[lang.Iso639One]app_msg_container.Container)

	for _, la := range priority {
		mc, err := app_msg_container_impl.New(la, c)
		if err != nil {
			return nil, false
		}
		containers[la.Code()] = mc
	}

	return wc.With(app_msg_container_impl.NewMultilingual(priority, containers)), true
}
