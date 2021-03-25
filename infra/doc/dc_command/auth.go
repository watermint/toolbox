package dc_command

import (
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
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
	demo := app_ui.MakeConsoleDemo(ui.Messages(), func(cui app_ui.UI) {
		rc_group.AppHeader(cui, "xx.x.xxx")
		cui.Info(api_auth_impl.MApiAuth.OauthSeq1.With("Url", "https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx"))
		cui.Info(api_auth_impl.MApiAuth.OauthSeq2)
	})
	ui.Info(z.ManualAuthDesc.With("Service", "Dropbox"))
	ui.Code(demo)
}
