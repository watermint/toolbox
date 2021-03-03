package es_filesystem_copier

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"io"
	"os"
	"time"
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

func (z localToLocalCopier) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return nil
}

func (z localToLocalCopier) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.l.With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")

	r, osErr := os.Open(source.Path().Path())
	if osErr != nil {
		l.Debug("Unable to open a file on the source path", esl.Error(osErr))
		onFailure(es_filesystem.NewCopyPair(source, target), es_filesystem_local.NewError(osErr))
		return
	}
	defer func() {
		_ = r.Close()
	}()

	w, osErr := os.Create(target.Path())
	if osErr != nil {
		l.Debug("Unable to create a file on the target path", esl.Error(osErr))
		onFailure(es_filesystem.NewCopyPair(source, target), es_filesystem_local.NewError(osErr))
		return
	}

	_, osErr = io.Copy(w, r)
	if osErr != nil {
		l.Debug("Unable to copy content, remove it", esl.Error(osErr))
		_ = w.Close()
		_ = os.Remove(target.Path())
		onFailure(es_filesystem.NewCopyPair(source, target), es_filesystem_local.NewError(osErr))
		return
	}

	_ = w.Close()

	osErr = os.Chtimes(target.Path(), time.Now(), source.ModTime())
	if osErr != nil {
		l.Debug("Unable to modify time", esl.Error(osErr))
		onFailure(es_filesystem.NewCopyPair(source, target), es_filesystem_local.NewError(osErr))
		return
	}

	if entry, err := z.target.Info(target); err != nil {
		l.Debug("Unable to resolve", esl.Error(err))
		onFailure(es_filesystem.NewCopyPair(source, target), err)
	} else {
		onSuccess(es_filesystem.NewCopyPair(source, target), entry)
	}
}

func (z localToLocalCopier) Shutdown() (err es_filesystem.FileSystemError) {
	z.l.Debug("Shutdown")
	return nil
}
