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
	if sel := Base(Select("en")); sel != "en" {
		t.Error(sel)
	}
	if sel := Base(Select("en-US")); sel != "en" {
		t.Error(sel)
	}

	if sel := Base(Select("ja-Jpan-JP")); sel != "ja" {
		t.Error(sel)
	}
	if sel := Base(Select("ja-JP")); sel != "ja" {
		t.Error(sel)
	}
	if sel := Base(Select("ja")); sel != "ja" {
		t.Error(sel)
	}

	// fallback to English
	if sel := Base(Select("und")); sel != "en" {
		t.Error(sel)
	}
}

func TestDetect(t *testing.T) {
	d := Detect()
	for _, l := range SupportedLanguages {
		if Base(l) == Base(d) {
			return
		}
	}
	t.Error(d)
}
