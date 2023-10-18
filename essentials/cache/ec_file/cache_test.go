package ec_file

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"strings"
	"testing"
)

func TestFileImpl_Get(t *testing.T) {
	cachePath, err := os.MkdirTemp("", "cache-file")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(cachePath)
	}()

	fc := New(cachePath, esl.Default())
	cacheFilePath, err := fc.Get("test", "read/me.txt", "https://raw.githubusercontent.com/watermint/toolbox/master/README.md")
	if err != nil {
		t.Error(err)
		return
	}
	if _, err := os.Lstat(cacheFilePath); err != nil {
		t.Error(err)
		return
	}
	cacheFileContent, err := os.ReadFile(cacheFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	if len(cacheFileContent) == 0 {
		t.Error("empty file")
		return
	}
	if !strings.Contains(string(cacheFileContent), "watermint toolbox") {
		t.Error("invalid file content")
		return
	}
}
