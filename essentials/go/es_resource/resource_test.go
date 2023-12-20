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
	r := NewNonTraversableResource("test", resTest)
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

func TestEmbedResource_Bytes(t *testing.T) {
	r := NewEmbedResource(map[string][]byte{
		"message.txt": []byte("hello"),
	})
	if x, err := r.Bytes("message.txt"); err != nil {
		t.Error(x, err)
	} else if string(x) != "hello" {
		t.Error(x, err)
	}
	if x, err := r.Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
}

func TestMergedResource_Bytes(t *testing.T) {
	r := NewMergedResource(
		NewEmbedResource(map[string][]byte{
			"message.txt": []byte("hello"),
		}),
		NewEmbedResource(map[string][]byte{
			"message.txt": []byte("world"),
			"hello.txt":   []byte("hello"),
		}),
	)
	if x, err := r.Bytes("message.txt"); err != nil {
		t.Error(x, err)
	} else if string(x) != "hello" {
		t.Error(x, err)
	}
	if x, err := r.Bytes("no existent"); err == nil {
		t.Error(x, err)
	}
	if x, err := r.Bytes("hello.txt"); err != nil {
		t.Error(x, err)
	} else if string(x) != "hello" {
		t.Error(x, err)
	}
}
