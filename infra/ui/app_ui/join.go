package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

func Join(ui UI, messages ...app_msg.Message) app_msg.Message {
	texts := make([]string, 0)
	for _, m := range messages {
		texts = append(texts, ui.Text(m))
	}
	return app_msg.Raw(strings.Join(texts, " "))
}
