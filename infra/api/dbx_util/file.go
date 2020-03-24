package dbx_util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	contentHashBlockSize = 4194304
	contentHashZeroHash  = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
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

// Calculate File content hash
func ContentHash(path string) (string, error) {
	l := app_root.Log().With(zap.String("path", path))

	info, err := os.Lstat(path)
	if err != nil {
		l.Debug("Unable to acquire information", zap.Error(err))
		return "", err
	}
	if info.Size() == 0 {
		return contentHashZeroHash, nil
	}

	f, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open file", zap.Error(err))
		return "", err
	}
	defer f.Close()

	var loadedBytes, totalBytes int64

	loadedBytes = 0
	totalBytes = info.Size()
	hashPerBlock := make([][32]byte, 0)

	for (totalBytes - loadedBytes) > 0 {
		r := io.LimitReader(f, contentHashBlockSize)
		block := make([]byte, contentHashBlockSize)
		readBytes, err := r.Read(block)
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Debug("unable to load file", zap.Error(err))
			return "", err
		}

		h := sha256.Sum256(block[:readBytes])
		hashPerBlock = append(hashPerBlock, h)
	}

	concatenated := make([]byte, 0)
	for _, h := range hashPerBlock {
		concatenated = append(concatenated, h[:]...)
	}
	h := sha256.Sum256(concatenated)
	return hex.EncodeToString(h[:]), nil
}
