package es_zip

import (
	"archive/zip"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrorUnsupportedOperation = errors.New("unsupported operation")
)

type ZipWriter interface {
	// Create an archive file. dstPath must include file name.
	Open(dstPath string) error

	// Add a file into the archive.
	AddFile(srcFilePath string, relPath string) error

	// Flush and close the archive file
	Close() error
}

func NewWriter(logger esl.Logger) ZipWriter {
	return &zwImpl{
		logger: logger,
	}
}

type zwImpl struct {
	logger  esl.Logger
	dstPath string
	w       *zip.Writer
	f       *os.File
}

func (z *zwImpl) Open(dstPath string) (err error) {
	l := z.logger.With(esl.String("destPath", dstPath))
	l.Debug("Create an archive file")
	z.f, err = os.Create(dstPath)
	if err != nil {
		l.Debug("Unable to create the file", esl.Error(err))
		return err
	}

	z.w = zip.NewWriter(z.f)
	z.dstPath = dstPath
	return nil
}

func (z *zwImpl) AddFile(srcFilePath string, relPath string) error {
	l := z.logger.With(esl.String("dstPath", z.dstPath), esl.String("srcFilePath", srcFilePath), esl.String("relPath", relPath))

	info, err := os.Lstat(srcFilePath)
	if err != nil {
		l.Debug("Unable to retrieve file info", esl.Error(err))
		return err
	}

	if info.IsDir() {
		l.Debug("Adding folder is not supported", esl.Any("info", info))
		return ErrorUnsupportedOperation
	}

	fr, err := os.Open(srcFilePath)
	if err != nil {
		l.Debug("Unable to read the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = fr.Close()
	}()

	fn := filepath.Join(relPath, info.Name())
	fw, err := z.w.CreateHeader(&zip.FileHeader{
		Name:     fn,
		Modified: info.ModTime(),
	})
	if err != nil {
		l.Debug("Unable to create the file entry", esl.Error(err))
		return err
	}

	size, err := io.Copy(fw, fr)
	if err != nil {
		l.Debug("Unable to add the file into the archive", esl.Error(err))
		return err
	}
	l.Debug("The file added", esl.Int64("size", size))

	return nil
}

func (z *zwImpl) Close() error {
	l := z.logger.With(esl.String("dstPath", z.dstPath))
	l.Debug("Close the archive")
	var lastErr error
	if err := z.w.Close(); err != nil {
		l.Debug("Unable to close the archive", esl.Error(err))
		lastErr = err
	}

	if err := z.f.Close(); err != nil {
		l.Debug("Unable to close the file", esl.Error(err))
		return err
	}
	return lastErr
}
