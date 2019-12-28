package ut_doc

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"text/template"
)

func msgFuncMap(ctl app_control.Control, test bool) template.FuncMap {
	return template.FuncMap{
		"msg": func(key string) string {
			if test {
				if !ctl.Messages().Exists(key) {
					ctl.UI().ErrorK(key)
				}
			}
			return ctl.Messages().Text(key)
		},
	}
}
