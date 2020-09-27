package es_file_copy

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
)

// Copy file from src path to dst path.
func Copy(src, dst string) error {
	l := esl.Default().With(esl.String("src", src), esl.String("dst", dst))
	l.Debug("Copy")

	srcFile, err := os.Open(src)
	if err != nil {
		l.Debug("Unable to open srcFile", esl.Error(err))
		return err
	}
	defer func() {
		_ = srcFile.Close()
	}()

	dstFile, err := os.Create(dst)
	if err != nil {
		l.Debug("Unable to create dstFile", esl.Error(err))
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()

	size, err := io.Copy(dstFile, srcFile)
	if err != nil {
		l.Debug("Unable to copy", esl.Error(err))
		return err
	}

	l.Debug("Copy completed", esl.Int64("size", size))
	return nil
}
