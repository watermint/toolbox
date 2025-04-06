package es_file_copy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/watermint/toolbox/quality/infra/qt_file"
)

func TestCopy(t *testing.T) {
	qt_file.TestWithTestFolder(t, "copy", false, func(path string) {
		srcPath := filepath.Join(path, "copy-src.txt")
		err := os.WriteFile(srcPath, []byte("0123abc"), 0644)
		if err != nil {
			t.Error(err)
			return
		}

		dstPath := filepath.Join(path, "copy-dst.txt")
		err = Copy(srcPath, dstPath)
		if err != nil {
			t.Error(err)
			return
		}

		dstContent, err := os.ReadFile(dstPath)
		if err != nil {
			t.Error(err)
		}

		if string(dstContent) != "0123abc" {
			t.Error(dstContent)
		}
	})
}
