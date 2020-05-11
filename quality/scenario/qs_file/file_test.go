package qs_file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScenario_CreateShort(t *testing.T) {
	sc, err := NewScenario(true)
	if err != nil {
		t.Error(err)
	}
	if sc.LocalPath == "" {
		t.Error(err)
	}
	if len(sc.Files) != 2 {
		t.Error(err)
	}
	if len(sc.Ignore) != 1 {
		t.Error(err)
	}

	for f := range sc.Files {
		_, err := os.Lstat(filepath.Join(sc.LocalPath, f))
		if err != nil {
			t.Error(err)
		}
	}
	for f := range sc.Ignore {
		_, err := os.Lstat(filepath.Join(sc.LocalPath, f))
		if err != nil {
			t.Error(err)
		}
	}

	if err := os.RemoveAll(sc.LocalPath); err != nil {
		t.Error(err)
	}
}
