package app_ui

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"testing"
)

func TestJoin(t *testing.T) {
	m1 := app_msg.CreateMessage("noodle", app_msg.P{"Message": "Ramen"})
	m2 := app_msg.CreateMessage("noodle", app_msg.P{"Message": "Udon"})
	mc := app_msg_container_impl.NewSingleWithMessages(map[string]string{
		"noodle": "Noodle[{{.Message}}]",
		"raw":    "{{.Raw}}",
	})
	ui := NewDiscard(mc, esl.Default())
	mj := Join(ui, m1, m2)
	if x := ui.Text(mj); x != "Noodle[Ramen] Noodle[Udon]" {
		t.Error(x)
	}
}
