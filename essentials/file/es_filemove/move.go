package es_filemove

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
)

// CopyThenDelete copies file into path `dst`, then deletes `src` file upon success.
func CopyThenDelete(src, dst string) error {
	l := esl.Default().With(esl.String("src", src), esl.String("dst", dst))

	srcInfo, err := os.Lstat(src)
	if err != nil {
		l.Debug("src info", esl.Error(err))
		return err
	}

	sf, err := os.Open(src)
	if err != nil {
		l.Debug("unable to open src", esl.Error(err))
		return err
	}
	defer func() {
		_ = sf.Close()
	}()

	df, err := os.Create(dst)
	if err != nil {
		l.Debug("unable to create dst", esl.Error(err))
		return err
	}

	cleanDstOnErr := func() {
		l.Debug("clean up dst")
		if df != nil {
			_ = df.Close()
		}
		_ = os.Remove(dst)
	}

	n, err := io.Copy(df, sf)
	if err != nil {
		l.Debug("unable to copy", esl.Error(err))
		cleanDstOnErr()
		return err
	}
	l.Debug("copy", esl.Int64("written", n))

	// flush
	err = df.Sync()
	if err != nil {
		l.Debug("unable to flush", esl.Error(err))
		cleanDstOnErr()
		return err
	}

	_ = df.Close()
	df = nil

	// copy file mode
	err = os.Chmod(dst, srcInfo.Mode())
	if err != nil {
		l.Debug("unable to update mode", esl.Error(err))
		cleanDstOnErr()
		return err
	}

	// copy file time
	err = os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime())
	if err != nil {
		l.Debug("unable to update mode", esl.Error(err))
		cleanDstOnErr()
		return err
	}

	// remove src
	err = os.Remove(src)
	l.Debug("remove src", esl.Error(err))
	return err
}

// Move moves file from path `src` to `dst`.
// Note: This function does not support moving folder.
func Move(src, dst string) error {
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return err
	}
	return CopyThenDelete(src, dst)
}
