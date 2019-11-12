package api_util

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

func IsFileNameIgnored(filename string) bool {
	// https://help.dropbox.com/installs-integrations/sync-uploads/files-not-syncing
	base := filepath.Base(filename)

	// file/folder name matching
	switch strings.ToLower(base) {
	case "desktop.ini", "thumbs.db", ".ds_store", "icon\r", ".dropbox", "dropbox.attr":
		return true
	}

	// patterns
	switch {
	case strings.HasPrefix(base, "~$"), strings.HasPrefix(base, ".~"):
		return true
	case strings.HasPrefix(base, "~") && strings.HasSuffix(base, ".tmp"):
		return true
	case strings.HasSuffix(base, "."):
		return true
	}

	// file/folder types
	stat, err := os.Lstat(filename)
	if err != nil {
		app_root.Log().Debug("unable to ensure file stat", zap.String("file", filename), zap.Error(err))
		return false
	}
	if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}

	return false
}
