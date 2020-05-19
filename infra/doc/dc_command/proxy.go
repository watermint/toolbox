package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewProxy(spec rc_recipe.Spec) dc_section.Section {
	return &Proxy{
		spec: spec,
	}
}

type Proxy struct {
	spec   rc_recipe.Spec
	Header app_msg.Message
	Desc   app_msg.Message
}

func (z Proxy) Title() app_msg.Message {
	return z.Header
}

func (z Proxy) Body(ui app_ui.UI) {
	ui.Info(z.Desc)
}
