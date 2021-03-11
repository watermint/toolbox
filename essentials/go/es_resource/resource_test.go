package es_resource

import (
	"embed"
	"testing"
)

//go:embed *_test.go
var resTest embed.FS

func TestNewResource(t *testing.T) {
	r := NewResource("test", resTest)
	if _, err := r.Bytes("no existent"); err == nil {
		t.Error(err)
	}
	hfs := r.HttpFileSystem()
	if _, err := hfs.Open("no existent"); err == nil {
		t.Error(err)
	}
}

func TestNewSecureResource(t *testing.T) {
	r := NewSecureResource("test", resTest)
	if _, err := r.Bytes("no existent"); err == nil {
		t.Error(err)
	}
	hfs := r.HttpFileSystem()
	if _, err := hfs.Open("no existent"); err == nil {
		t.Error(err)
	}
}

func TestNewEmptyResource(t *testing.T) {
	r := EmptyResource()
	if _, err := r.Bytes("no existent"); err == nil {
		t.Error(err)
	}
	hfs := r.HttpFileSystem()
	if _, err := hfs.Open("no existent"); err == nil {
		t.Error(err)
	}
}

func TestEmptyBundle(t *testing.T) {
	b := EmptyBundle()

	if x, err := b.Templates().Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
	if x, err := b.Messages().Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
	if x, err := b.Web().Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
	if x, err := b.Keys().Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
	if x, err := b.Images().Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
	if x, err := b.Data().Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
}
