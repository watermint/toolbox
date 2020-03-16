package app_lang

import (
	"golang.org/x/text/language"
	"testing"
)

func TestBase(t *testing.T) {
	if base := Base(language.English); base != "en" {
		t.Error(base)
	}
	if base := Base(language.Japanese); base != "ja" {
		t.Error(base)
	}
}

func TestPathSuffix(t *testing.T) {
	if suffix := PathSuffix(language.English); suffix != "" {
		t.Error(suffix)
	}
	if suffix := PathSuffix(language.Japanese); suffix != "_ja" {
		t.Error(suffix)
	}
}

func TestSelect(t *testing.T) {
	if sel := Select("en"); sel != language.English {
		t.Error(sel)
	}
	if sel := Select("en-US"); sel != language.English {
		t.Error(sel)
	}

	if sel := Select("ja-Jpan-JP"); sel != language.Japanese {
		t.Error(sel)
	}
	if sel := Select("ja-JP"); sel != language.Japanese {
		t.Error(sel)
	}
	if sel := Select("ja"); sel != language.Japanese {
		t.Error(sel)
	}

	// fallback to English
	if sel := Select("und"); sel != language.English {
		t.Error(sel)
	}
}

func TestDetect(t *testing.T) {
	d := Detect()
	for _, l := range SupportedLanguages {
		if l == d {
			return
		}
	}
	t.Error(d)
}
