package es_filepath

import (
	"path/filepath"
	"strings"
)

func IsSystemFile(path string) bool {
	base := filepath.Base(path)
	if strings.HasPrefix(base, ".") {
		return true
	}

	switch strings.ToLower(base) {
	case "desktop.ini", "thumbs.db", ".ds_store", "icon\r", ".dropbox", ".dropbox.attr":
		return true
	default:
		return false
	}
}
