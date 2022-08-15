package es_filemove

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"io/fs"
	"os"
)

var (
	ErrCannotMoveDir        = errors.New("this function does not support folder move")
	ErrCannotMoveAppend     = errors.New("this function does not support append only file")
	ErrCannotMoveExclusive  = errors.New("this function does not support exclusive file")
	ErrCannotMoveTemporary  = errors.New("this function does not support temporary file")
	ErrCannotMoveDevice     = errors.New("this function does not support device file")
	ErrCannotMoveNamedPipe  = errors.New("this function does not support named pipe (FIFO)")
	ErrCannotMoveSocket     = errors.New("this function does not support Unix domain socket")
	ErrCannotMoveSetuid     = errors.New("this function does not support setuid file")
	ErrCannotMoveSetgid     = errors.New("this function does not support setgid file")
	ErrCannotMoveCharDevice = errors.New("this function does not support Unix character device")
	ErrCannotMoveSticky     = errors.New("this function does not support sticky")
	ErrCannotMoveIrregular  = errors.New("this function does not support non-regular file")
)

// copySymLinkThenDelete copies symlink to dst.
func copySymLinkThenDelete(src string, dst string) error {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}

	err = os.Symlink(dst, link)
	if err != nil {
		return err
	}

	return os.Remove(src)
}

// CopyThenDelete copies file into path `dst`, then deletes `src` file upon success.
// The src file may be kept in case of error during move.
func CopyThenDelete(src, dst string) error {
	l := esl.Default().With(esl.String("src", src), esl.String("dst", dst))

	srcInfo, err := os.Lstat(src)
	if err != nil {
		l.Debug("src info", esl.Error(err))
		return err
	}

	// handle special files
	switch {
	case srcInfo.IsDir():
		return ErrCannotMoveDir
	case srcInfo.Mode()&fs.ModeSymlink != 0:
		return copySymLinkThenDelete(src, dst)
	case srcInfo.Mode()&fs.ModeAppend != 0:
		return ErrCannotMoveAppend
	case srcInfo.Mode()&fs.ModeExclusive != 0:
		return ErrCannotMoveExclusive
	case srcInfo.Mode()&fs.ModeTemporary != 0:
		return ErrCannotMoveTemporary
	case srcInfo.Mode()&fs.ModeDevice != 0:
		return ErrCannotMoveDevice
	case srcInfo.Mode()&fs.ModeNamedPipe != 0:
		return ErrCannotMoveNamedPipe
	case srcInfo.Mode()&fs.ModeSocket != 0:
		return ErrCannotMoveSocket
	case srcInfo.Mode()&fs.ModeSetuid != 0:
		return ErrCannotMoveSetuid
	case srcInfo.Mode()&fs.ModeSetgid != 0:
		return ErrCannotMoveSetgid
	case srcInfo.Mode()&fs.ModeCharDevice != 0:
		return ErrCannotMoveCharDevice
	case srcInfo.Mode()&fs.ModeSticky != 0:
		return ErrCannotMoveSticky
	case srcInfo.Mode()&fs.ModeIrregular != 0:
		return ErrCannotMoveIrregular
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
// The src file may be kept in case of error during move.
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
