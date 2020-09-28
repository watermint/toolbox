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
	ErrorConflict = errors.New("conflict")
)

// Extract the archive file to dest path.
func Extract(log esl.Logger, archivePath, destPath string) error {
	l := log.With(esl.String("archivePath", archivePath), esl.String("destPath", destPath))
	destAbsPath, err := filepath.Abs(destPath)
	if err != nil {
		l.Debug("Unable to compute abs dest path", esl.Error(err))
		return err
	}

	zr, err := zip.OpenReader(archivePath)
	if err != nil {
		l.Debug("Unable to open the archive")
		return err
	}
	defer func() {
		_ = zr.Close()
	}()

	extractFile := func(zf *zip.File) error {
		fileFolder := filepath.Join(destAbsPath, filepath.Dir(zf.Name))
		ll := l.With(esl.String("fileFolder", fileFolder), esl.String("filePath", zf.Name))

		zr, err := zf.Open()
		if err != nil {
			l.Debug("Unable to read the file", esl.Error(err))
			return err
		}
		defer func() {
			_ = zr.Close()
		}()

		fileFolderInfo, err := os.Lstat(fileFolder)
		switch {
		case os.IsNotExist(err):
			ll.Debug("Try create a folder")
			if err = os.MkdirAll(fileFolder, 0755); err != nil {
				ll.Debug("Unable to create the folder", esl.Error(err))
				return err
			}
		case fileFolderInfo != nil && !fileFolderInfo.IsDir():
			ll.Debug("Path conflict with a file")
			return ErrorConflict
		case err == nil:
			ll.Debug("The folder found")
		default:
			ll.Debug("Unable to determine the folder", esl.Error(err))
			return err
		}

		filePath := filepath.Join(destAbsPath, zf.Name)
		f, err := os.Create(filePath)
		if err != nil {
			ll.Debug("Unable to create the file", esl.Error(err))
			return err
		}
		defer func() {
			_ = f.Close()
		}()

		size, err := io.Copy(f, zr)
		if err != nil {
			ll.Debug("Unable to copy", esl.Error(err))
			return err
		}
		ll.Debug("Extract completed", esl.Int64("size", size))
		return nil
	}

	for _, zf := range zr.File {
		if err := extractFile(zf); err != nil {
			return err
		}
	}
	return nil
}
