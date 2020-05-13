package app_msg_container_impl

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"testing"
)

func TestAltCompile(t *testing.T) {
	amc := &Alt{}

	// always false
	if x := amc.ExistsKey("ping"); x {
		t.Error(x)
	}

	if x := amc.Text("ping"); x != "Key[ping]" {
		t.Error(x)
	}
	m := app_msg.CreateMessage("ping", app_msg.P{"Pong": "Pang"})
	if x := amc.Compile(m); x != `{"key":"ping","params":{"Pong":"Pang"}}` {
		t.Error(x)
	}
}
