package lang

import (
	"golang.org/x/text/language"
	"testing"
)

func TestSelect(t *testing.T) {
	// fallback
	if s := Select("", []Lang{Japanese, English}); s.Code() != "en" {
		t.Error(s.Code())
	}

	// fallback
	if s := Select("de", []Lang{Japanese, English}); s.Code() != "en" {
		t.Error(s.Code())
	}

	// exact match
	if s := Select("ja", []Lang{Japanese, English}); s.Code() != "ja" {
		t.Error(s.Code())
	}

	// no confidence
	if s := Select("zh", []Lang{Japanese, English}); s.Code() != "en" {
		t.Error(s.Code())
	}

	// low confidence
	if s := Select("zh_CN", []Lang{English, New(language.TraditionalChinese)}); s.Code() != "en" {
		t.Error(s.Code())
	}
}
