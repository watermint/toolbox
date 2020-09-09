package dbx_util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/watermint/toolbox/essentials/log/esl"
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
		esl.Default().Debug("unable to ensure file stat", esl.String("file", filename), esl.Error(err))
		return false
	}
	if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}

	return false
}

func ContentHash(f io.ReadCloser, totalBytes int64) (string, error) {
	l := esl.Default().With(esl.Int64("totalBytes", totalBytes))
	var loadedBytes int64

	loadedBytes = 0
	hashPerBlock := make([][32]byte, 0)

	for (totalBytes - loadedBytes) > 0 {
		r := io.LimitReader(f, contentHashBlockSize)
		block := make([]byte, contentHashBlockSize)
		readBytes, err := r.Read(block)
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Debug("unable to load data", esl.Error(err))
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

// Calculate File content hash
func FileContentHash(path string) (string, error) {
	l := esl.Default().With(esl.String("path", path))

	info, err := os.Lstat(path)
	if err != nil {
		l.Debug("Unable to acquire information", esl.Error(err))
		return "", err
	}
	if info.Size() == 0 {
		return contentHashZeroHash, nil
	}

	f, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open file", esl.Error(err))
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	return ContentHash(f, info.Size())
}
