package dc_command

import (
	"github.com/watermint/toolbox/essentials/api/api_doc"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewAuth(spec rc_recipe.Spec) dc_section.Section {
	return &Auth{
		spec: spec,
	}
}

type Auth struct {
	spec           rc_recipe.Spec
	Header         app_msg.Message
	ManualAuthDesc app_msg.Message
}

func (z Auth) Title() app_msg.Message {
	return z.Header
}

func (z Auth) Body(ui app_ui.UI) {
	services := z.spec.Services()
	msgBase := es_reflect.Key(app.Pkg, &z)
	for _, service := range services {
		serviceName := ui.Text(app_msg.CreateMessage(msgBase + ".service_name." + service))
		serviceCuiPreview, ok := api_doc.ApiDocCuiPreview[service]
		if !ok {
			panic("No Api Auth CUI preview document found for the service [" + service + "]")
		}
		serviceAuthDesc, ok := api_doc.ApiDocAuthDesc[service]
		if !ok {
			panic("No Api Auth desc document found for the service [" + service + "]")
		}

		demo := app_ui.MakeConsoleDemo(ui.Messages(), func(cui app_ui.UI) {
			rc_group.AppHeader(cui, "xx.x.xxx")
			serviceCuiPreview(cui)
		})
		ui.Info(z.ManualAuthDesc.With("Service", serviceName))
		ui.Info(serviceAuthDesc)
		ui.Code(demo)
	}
}
