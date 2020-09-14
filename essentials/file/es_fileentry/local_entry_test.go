package es_fileentry

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestReadLocalEntries(t *testing.T) {
	wd, err := filepath.Abs(".")
	if err != nil {
		t.Error(err)
		return
	}
	localEntries, err := ReadLocalEntries(wd)
	if err != nil {
		t.Error(err)
	}

	for _, localEntry := range localEntries {
		if !strings.HasPrefix(localEntry.Path, wd) {
			t.Error(localEntry)
		}
	}

	byName := LocalEntryByNameLower(localEntries)
	for name, localEntry := range byName {
		if name != strings.ToLower(localEntry.Name) {
			t.Error(localEntry)
		}
	}
}
