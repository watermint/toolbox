package es_gzip

import (
	"compress/gzip"
	"io"
	"os"
	"testing"
	"time"

	"github.com/watermint/toolbox/quality/infra/qt_file"
)

func TestExecCompress(t *testing.T) {
	content := time.Now().String()
	qt_file.TestWithTestFile(t, "test", content, func(path string) {
		if _, err := Compress(path); err != nil {
			t.Error(err)
		}

		// should not exist
		fi, err := os.Lstat(path)
		if err == nil {
			t.Error(fi, err)
		}

		cp := path + SuffixCompress
		ci, err := os.Lstat(cp)
		if err != nil {
			t.Error(ci, err)
		}

		cc, err := os.Open(cp)
		if err != nil {
			t.Error(err)
		}

		cr, err := gzip.NewReader(cc)
		if err != nil {
			t.Error(err)
		}
		dc, err := io.ReadAll(cr)
		if err != nil {
			t.Error(err)
		}
		_ = cc.Close()
		if string(dc) != content {
			t.Error(dc, content)
		}
	})
}
