package app_msg_container_impl

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"testing"
)

func TestSglContainer_Text(t *testing.T) {
	sgl := NewSingleWithMessagesForTest(map[string]string{
		"ping":   "Pong",
		"noodle": "Ramen",
	})

	if x := sgl.Text("ping"); x != "Pong" {
		t.Error(x)
	}
	if x := sgl.ExistsKey("ping"); !x {
		t.Error(x)
	}
	if x := sgl.ExistsKey("pong"); x {
		t.Error(x)
	}
	m1 := app_msg.CreateMessage("ping")
	if x := sgl.Compile(m1); x != "Pong" {
		t.Error(x)
	}

	m2 := app_msg.CreateMessage("noodle")
	mj := app_msg.Join(m1, m2)
	if x := sgl.Compile(mj); x != "Pong Ramen" {
		t.Error(x)
	}

	m3 := app_msg.CreateMessage("steak").AsOptional()
	if x := sgl.Compile(m3); x != "" {
		t.Error(x)
	}
}
