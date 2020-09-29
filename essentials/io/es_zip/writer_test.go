package es_zip

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestZwImpl_AddFile(t *testing.T) {
	basePath, err := ioutil.TempDir("", "zip")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(basePath)
	}()
	testFilePath1 := filepath.Join(basePath, "hello.txt")
	testFilePath2 := filepath.Join(basePath, "thank_you.txt")
	archivePath := filepath.Join(basePath, "test.zip")

	if err = ioutil.WriteFile(testFilePath1, []byte("hello"), 0644); err != nil {
		t.Error(err)
		return
	}
	if err = ioutil.WriteFile(testFilePath2, []byte("thank you"), 0644); err != nil {
		t.Error(err)
		return
	}

	zw := NewWriter(esl.Default())
	if err := zw.Open(archivePath); err != nil {
		t.Error(err)
	}

	if err := zw.AddFile(testFilePath1, "greetings"); err != nil {
		t.Error(err)
	}
	if err := zw.AddFile(testFilePath2, "greetings"); err != nil {
		t.Error(err)
	}

	if err := zw.Close(); err != nil {
		t.Error(err)
	}

	if err := Extract(esl.Default(), archivePath, filepath.Join(basePath, "extract")); err != nil {
		t.Error(err)
	}

	extractFile1, err := ioutil.ReadFile(filepath.Join(basePath, "extract", "greetings", "hello.txt"))
	if err != nil {
		t.Error(err)
	}

	if string(extractFile1) != "hello" {
		t.Error(err)
	}

	extractFile2, err := ioutil.ReadFile(filepath.Join(basePath, "extract", "greetings", "thank_you.txt"))
	if err != nil {
		t.Error(err)
	}

	if string(extractFile2) != "thank you" {
		t.Error(err)
	}
}
