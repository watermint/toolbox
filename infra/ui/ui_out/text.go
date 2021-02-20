package ui_out

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func TextOut(c app_control.Control, text string) {
	if c.Feature().IsQuiet() {
		_, _ = es_stdout.NewDirectOut().Write([]byte(text))
	} else {
		c.UI().Info(app_msg.Raw(text))
	}
}
