package compare

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/cihub/seelog"
	"io"
	"os"
)

const (
	BLOCK_SIZE     = 4194304
	HASH_FOR_EMPTY = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

// Calculate File content hash
func ContentHash(path string) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] error[%s]", path, err)
		return "", err
	}

	f, err := os.Open(path)
	if err != nil {
		seelog.Warnf("Unable to open file [%s] error[%s]", path, err)
		return "", err
	}
	defer f.Close()

	var loadedBytes, totalBytes int64

	loadedBytes = 0
	totalBytes = info.Size()
	hashePerBlock := make([][32]byte, 0)

	for (totalBytes - loadedBytes) > 0 {
		r := io.LimitReader(f, BLOCK_SIZE)
		block := make([]byte, BLOCK_SIZE)
		readBytes, err := r.Read(block)
		if err == io.EOF {
			break
		}
		if err != nil {
			seelog.Warnf("Unable to load file [%s] error[%s]", path, err)
			return "", err
		}

		h := sha256.Sum256(block[:readBytes])
		hashePerBlock = append(hashePerBlock, h)
	}

	if len(hashePerBlock) < 1 {
		return HASH_FOR_EMPTY, nil
	} else {
		concatenated := make([]byte, 0)
		for _, h := range hashePerBlock {
			concatenated = append(concatenated, h[:]...)
		}
		h := sha256.Sum256(concatenated)
		return hex.EncodeToString(h[:]), nil
	}
}
