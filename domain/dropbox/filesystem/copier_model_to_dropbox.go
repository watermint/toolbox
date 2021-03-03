package filesystem

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func NewModelToDropbox(l esl.Logger, modelRoot em_file.Folder, conn es_filesystem.Connector) es_filesystem.Connector {
	return &copierModelToDropbox{
		l:         l,
		conn:      conn,
		modelRoot: modelRoot,
	}
}

type copierModelToDropbox struct {
	l         esl.Logger
	conn      es_filesystem.Connector
	modelRoot em_file.Folder
}

func (z copierModelToDropbox) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return z.conn.Startup(qd)
}

func (z copierModelToDropbox) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.l.With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")
	cp := es_filesystem.NewCopyPair(source, target)

	sourceNode := em_file.ResolvePath(z.modelRoot, source.Path().Path())
	if sourceNode == nil {
		l.Debug("Unable to find the source node")
		onFailure(cp, es_filesystem_model.NewError(errors.New("source node not found"), es_filesystem_model.ErrorTypePathNotFound))
		return
	}

	if sourceNode.Type() != em_file.FileNode {
		l.Debug("Node is not a file")
		onFailure(cp, es_filesystem_model.NewError(errors.New("source node is not a file"), es_filesystem_model.ErrorTypeOther))
		return
	}

	content := sourceNode.(em_file.File).Content()

	tmpDir, ioErr := ioutil.TempDir("", "model_to_dropbox")
	if ioErr != nil {
		l.Debug("unable to create temp file", esl.Error(ioErr))
		onFailure(cp, NewError(ioErr))
		return
	}

	tmpFilePath := filepath.Join(tmpDir, sourceNode.Name())

	if errIO := ioutil.WriteFile(tmpFilePath, content, 0644); errIO != nil {
		l.Debug("Unable to write to the file", esl.Error(errIO))
		onFailure(cp, es_filesystem_local.NewError(errIO))
		return
	}

	if errIO := os.Chtimes(tmpFilePath, time.Now(), source.ModTime()); errIO != nil {
		l.Debug("Unable to modify time", esl.Error(errIO))
		onFailure(cp, es_filesystem_local.NewError(errIO))
		return
	}

	tmpFileInfo, osErr := os.Lstat(tmpFilePath)
	if osErr != nil {
		l.Debug("Unable to retrieve file metadata", esl.Error(osErr))
		onFailure(cp, es_filesystem_local.NewError(osErr))
		return
	}

	z.conn.Copy(
		es_filesystem_local.NewEntry(tmpFilePath, tmpFileInfo), target,
		func(pair es_filesystem.CopyPair, copied es_filesystem.Entry) {
			onSuccess(pair, copied)
			_ = os.RemoveAll(tmpDir)
		},
		func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError) {
			onFailure(pair, err)
			_ = os.RemoveAll(tmpDir)
		},
	)
}

func (z copierModelToDropbox) Shutdown() (err es_filesystem.FileSystemError) {
	z.l.Debug("Shutdown")
	return z.conn.Shutdown()
}
