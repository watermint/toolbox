package dc_command

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"text/template"
)

func msgFuncMap(ctl app_control.Control) template.FuncMap {
	return template.FuncMap{
		"msg": func(key string) string {
			if ctl.Feature().IsTest() {
				if !ctl.Messages().ExistsKey(key) {
					ctl.UI().Error(app_msg.CreateMessage(key))
				}
			}
			return ctl.Messages().Text(key)
		},
	}
}
