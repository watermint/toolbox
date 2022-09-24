package es_lang

import (
	"golang.org/x/text/language"
	"testing"
)

func TestLangImpl_Code(t *testing.T) {
	if c := New(language.English).Code(); c != "en" {
		t.Error(c)
	}
	if c := New(language.Japanese).Code(); c != "ja" {
		t.Error(c)
	}
	if c := New(language.German).Code(); c != "de" {
		t.Error(c)
	}
	if c := New(language.Chinese).Code(); c != "zh" {
		t.Error(c)
	}
}

func TestLangImpl_Suffix(t *testing.T) {
	if c := New(language.English).Suffix(); c != "" {
		t.Error(c)
	}
	if c := New(language.Japanese).Suffix(); c != "_ja" {
		t.Error(c)
	}
}

func TestLangImpl_IsDefault(t *testing.T) {
	// high confidence
	if c := New(language.English).Equals(New(language.Make("en_UK"))); !c {
		t.Error(c)
	}

	// low confidence
	if c := New(language.Chinese).Equals(New(language.Make("zh_HK"))); c {
		t.Error(c)
	}
}

func TestLangImpl_Tag(t *testing.T) {
	if c := New(language.English).Tag(); c != language.English {
		t.Error(c)
	}
	if c := New(language.Japanese).Tag(); c != language.Japanese {
		t.Error(c)
	}
}
