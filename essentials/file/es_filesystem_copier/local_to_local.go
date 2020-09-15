package es_filesystem_copier

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
	"time"
)

var (
	ErrorUnsupportedOperation = errors.New("unsupported operation")
)

func NewLocalToLocal(l esl.Logger, source, target es_filesystem.FileSystem) es_filesystem.Connector {
	return &localToLocalCopier{
		l:      l,
		source: source,
		target: target,
	}
}

type localToLocalCopier struct {
	l      esl.Logger
	source es_filesystem.FileSystem
	target es_filesystem.FileSystem
}

func (z localToLocalCopier) Copy(source es_filesystem.Entry, target es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.l.With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")

	r, osErr := os.Open(source.Path().Path())
	if osErr != nil {
		l.Debug("Unable to open a file on the source path", esl.Error(err))
		return nil, es_filesystem_local.NewError(osErr)
	}
	defer func() {
		_ = r.Close()
	}()

	w, osErr := os.Create(target.Path())
	if osErr != nil {
		l.Debug("Unable to create a file on the target path", esl.Error(err))
		return nil, es_filesystem_local.NewError(osErr)
	}

	_, osErr = io.Copy(w, r)
	if osErr != nil {
		l.Debug("Unable to copy content, remove it", esl.Error(err))
		_ = w.Close()
		_ = os.Remove(target.Path())
		return nil, es_filesystem_local.NewError(osErr)
	}

	_ = w.Close()

	osErr = os.Chtimes(target.Path(), time.Now(), source.ModTime())
	if osErr != nil {
		l.Debug("Unable to modify time", esl.Error(err))
	}

	return z.target.Info(target)
}
