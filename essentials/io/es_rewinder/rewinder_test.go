package es_rewinder

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestNewReadRewinderJsonStruct(t *testing.T) {
	d := struct {
		SKU  string `json:"sku"`
		Name string `json:"name"`
	}{
		SKU:  "S100",
		Name: "Smartphone",
	}
	rr, err := NewReadRewinderJsonStruct(d)
	if err != nil {
		t.Error(err)
	}
	if rr.Length() < 1 {
		t.Error(rr.Length())
	}
}

func TestReaderRewinderWithLimit_Read(t *testing.T) {
	d, err := os.MkdirTemp("", "read_rewinder")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(d)
	}()
	path := filepath.Join(d, "readrewinder.dat")
	f, err := os.Create(path)
	if err != nil {
		t.Error(err)
		return
	}
	data := `0123456789abcdefghijABCDEFGHIJ`
	_, err = f.WriteString(data)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		t.Error(err)
		return
	}

	{
		lr, err := NewReadRewinderWithLimit(f, 0, 10)
		if err != nil {
			t.Error(err)
			return
		}
		if lr.Length() != 10 {
			t.Error(lr.Length())
		}
		b0_10, err := io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if string(b0_10) != "0123456789" {
			t.Error(b0_10)
			return
		}
	}

	{
		lr, err := NewReadRewinderWithLimit(f, 10, 10)
		if err != nil {
			t.Error(err)
			return
		}
		b10_20, err := io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if lr.Length() != 10 {
			t.Error(lr.Length())
		}
		if string(b10_20) != "abcdefghij" {
			t.Error(b10_20)
			return
		}
		lr.Rewind()
		b10_20, err = io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if string(b10_20) != "abcdefghij" {
			t.Error(b10_20)
			return
		}
	}

	{
		lr, err := NewReadRewinderWithLimit(f, 20, 10)
		if err != nil {
			t.Error(err)
			return
		}
		b20_30, err := io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if lr.Length() != 10 {
			t.Error(lr.Length())
		}
		if string(b20_30) != "ABCDEFGHIJ" {
			t.Error(b20_30)
			return
		}
		lr.Rewind()
		b20_30, err = io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if string(b20_30) != "ABCDEFGHIJ" {
			t.Error(b20_30)
			return
		}
	}

	{
		lr, err := NewReadRewinderWithLimit(f, 0, 10)
		if err != nil {
			t.Error(err)
			return
		}
		if lr.Length() != 10 {
			t.Error(lr.Length())
		}
		b0_10, err := io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if string(b0_10) != "0123456789" {
			t.Error(b0_10)
			return
		}
	}

	{
		lr, err := NewReadRewinderWithLimit(f, 25, 10)
		if err != nil {
			t.Error(err)
			return
		}
		b25_30, err := io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if lr.Length() != 5 {
			t.Error(lr.Length())
		}
		if string(b25_30) != "FGHIJ" {
			t.Error(b25_30)
			return
		}
		lr.Rewind()
		b25_30, err = io.ReadAll(lr)
		if err != nil {
			t.Error(err)
			return
		}
		if string(b25_30) != "FGHIJ" {
			t.Error(b25_30)
			return
		}
	}
}
