package app_msg_container_impl

import (
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"golang.org/x/text/language"
	"testing"
)

func TestNewMultilingual(t *testing.T) {
	en := &Resource{Messages: map[string]string{
		"ping": "Ping",
		"pong": "Pong",
	}}
	ja := &Resource{Messages: map[string]string{
		"ping": "ピン",
	}}
	containers := map[language.Tag]app_msg_container.Container{
		language.English:  en,
		language.Japanese: ja,
	}

	enJa := NewMultilingual(
		[]language.Tag{
			language.English,
			language.Japanese,
		},
		containers,
	)

	if x := enJa.Text("ping"); x != "Ping" {
		t.Error(x)
	}
	if x := enJa.Text("pong"); x != "Pong" {
		t.Error(x)
	}

	jaEn := NewMultilingual(
		[]language.Tag{
			language.Japanese,
			language.English,
		},
		containers,
	)

	if x := jaEn.Text("ping"); x != "ピン" {
		t.Error(x)
	}
	// should fallback
	if x := jaEn.Text("pong"); x != "Pong" {
		t.Error(x)
	}
}
