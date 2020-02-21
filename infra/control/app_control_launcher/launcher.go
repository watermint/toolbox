package app_control_launcher

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"golang.org/x/text/language"
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

func ControlWithLang(lang string, c app_control.Control) (app_control.Control, bool) {
	wc, ok := c.(WithMessageContainer)
	if !ok {
		return nil, false
	}

	langPriority := make([]language.Tag, 0)
	ul := app_lang.Select(lang)
	if ul != language.English {
		langPriority = append(langPriority, ul)
	}
	langPriority = append(langPriority, language.English)
	langContainers := make(map[language.Tag]app_msg_container.Container)

	for _, lang := range langPriority {
		mc, err := app_msg_container_impl.New(lang, c)
		if err != nil {
			return nil, false
		}
		langContainers[lang] = mc
	}

	return wc.With(app_msg_container_impl.NewMultilingual(langPriority, langContainers)), true
}
