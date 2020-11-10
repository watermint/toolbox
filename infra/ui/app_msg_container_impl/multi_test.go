package app_msg_container_impl

import (
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"testing"
)

func TestNewMultilingual(t *testing.T) {
	en := NewSingleWithMessagesForTest(map[string]string{
		"ping": "Ping",
		"pong": "Pong",
	})
	ja := NewSingleWithMessagesForTest(map[string]string{
		"ping": "ピン",
	})
	containers := map[lang.Iso639One]app_msg_container.Container{
		"en": en,
		"ja": ja,
	}

	enJa := NewMultilingual(
		[]lang.Lang{lang.English, lang.Japanese},
		containers,
	)

	if x := enJa.Text("ping"); x != "Ping" {
		t.Error(x)
	}
	if x := enJa.Text("pong"); x != "Pong" {
		t.Error(x)
	}

	jaEn := NewMultilingual(
		[]lang.Lang{lang.Japanese, lang.English},
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
