package es_filesystem_copier

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
	"time"
)

func NewModelToLocal(l esl.Logger, sourceRoot em_file.Folder) es_filesystem.Connector {
	return &modelToLocalCopier{
		l:          l,
		sourceRoot: sourceRoot,
		target:     es_filesystem_local.NewFileSystem(),
	}
}

type modelToLocalCopier struct {
	l          esl.Logger
	sourceRoot em_file.Folder
	target     es_filesystem.FileSystem
}

func (z modelToLocalCopier) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return nil
}

func (z modelToLocalCopier) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.l.With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")
	cp := es_filesystem.NewCopyPair(source, target)

	sourceNode := em_file.ResolvePath(z.sourceRoot, source.Path().Path())
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

	targetFolder := target.Ancestor().Path()
	targetFolderInfo, osErr := os.Lstat(targetFolder)
	switch {
	case osErr == nil:
		if !targetFolderInfo.IsDir() {
			l.Debug("Target folder path is not folder")
			onFailure(cp, es_filesystem_local.NewError(errors.New("target folder path is not a folder")))
			return
		}
	case os.IsNotExist(osErr):
		l.Debug("Target folder not found, create it", esl.Error(osErr))
		osErr = os.MkdirAll(targetFolder, 0755)
		if osErr != nil {
			l.Debug("Unable to create folder", esl.Error(osErr))
			onFailure(cp, es_filesystem_local.NewError(osErr))
		}
	default:
		l.Debug("Unable to retrieve target folder information", esl.Error(osErr))
		onFailure(cp, es_filesystem_local.NewError(osErr))
		return
	}

	content := sourceNode.(em_file.File).Content()

	if errIO := ioutil.WriteFile(target.Path(), content, 0644); errIO != nil {
		l.Debug("Unable to write to the file", esl.Error(errIO))
		onFailure(cp, es_filesystem_local.NewError(errIO))
		return
	}

	if errIO := os.Chtimes(target.Path(), time.Now(), source.ModTime()); errIO != nil {
		l.Debug("Unable to modify time", esl.Error(errIO))
		onFailure(cp, es_filesystem_local.NewError(errIO))
		return
	}

	entry, fsErr := z.target.Info(target)
	if fsErr != nil {
		l.Debug("Unable to resolve the entry", esl.Error(fsErr))
		onFailure(cp, fsErr)
	} else {
		onSuccess(cp, entry)
	}
}

func (z modelToLocalCopier) Shutdown() (err es_filesystem.FileSystemError) {
	z.l.Debug("Shutdown")
	return nil
}
