package es_filesystem_local

import (
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"path/filepath"
	"testing"
)

func TestNewFileSystem(t *testing.T) {
	qt_file.TestWithTestFolder(t, "fs", true, func(path string) {
		comparator := es_filecompare.New()

		fs := NewFileSystem()
		entries, err := fs.List(NewPath(path))
		if err != nil {
			t.Error(err)
		}
		if len(entries) < 1 {
			t.Error(entries)
		}

		for _, entry := range entries {
			resolvedEntry, err := fs.Info(entry.Path())
			if err != nil {
				t.Error(entry, err)
			}

			if same, err := comparator.Compare(entry, resolvedEntry); !same || err != nil {
				t.Error(entry, same, err)
			}

			if err := fs.Delete(entry.Path()); err != nil {
				t.Error(entry, err)
			}
		}

		if _, err := fs.CreateFolder(NewPath(filepath.Join(path, "hello"))); err != nil {
			t.Error(err)
		}
	})
}
