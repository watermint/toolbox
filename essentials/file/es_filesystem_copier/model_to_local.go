package es_filesystem_copier

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_tree"
	"io/ioutil"
	"os"
	"time"
)

func NewModelToLocal(l esl.Logger, sourceRoot em_tree.Folder) es_filesystem.Connector {
	return &modelToLocalCopier{
		l:          l,
		sourceRoot: sourceRoot,
		target:     es_filesystem_local.NewFileSystem(),
	}
}

type modelToLocalCopier struct {
	l          esl.Logger
	sourceRoot em_tree.Folder
	target     es_filesystem.FileSystem
}

func (z modelToLocalCopier) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.l.With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")

	sourceNode := em_tree.ResolvePath(z.sourceRoot, source.Path().Path())
	if sourceNode == nil {
		l.Debug("Unable to find the source node")
		return nil, es_filesystem_model.NewError(errors.New("source node not found"), es_filesystem_model.ErrorTypePathNotFound)
	}

	if sourceNode.Type() != em_tree.FileNode {
		l.Debug("Node is not a file")
		return nil, es_filesystem_model.NewError(errors.New("source node is not a file"), es_filesystem_model.ErrorTypeOther)
	}

	targetFolder := target.Ancestor().Path()
	targetFolderInfo, osErr := os.Lstat(targetFolder)
	switch {
	case osErr == nil:
		if !targetFolderInfo.IsDir() {
			l.Debug("Target folder path is not folder")
			return nil, es_filesystem_local.NewError(errors.New("target folder path is not a folder"))
		}
	case os.IsNotExist(osErr):
		l.Debug("Target folder not found, create it", esl.Error(osErr))
		osErr = os.MkdirAll(targetFolder, 0755)
		if osErr != nil {
			l.Debug("Unable to create folder", esl.Error(osErr))
			return nil, es_filesystem_local.NewError(osErr)
		}
	default:
		l.Debug("Unable to retrieve target folder information", esl.Error(osErr))
		return nil, es_filesystem_local.NewError(osErr)
	}

	content := sourceNode.(em_tree.File).Content()

	if errIO := ioutil.WriteFile(target.Path(), content, 0644); errIO != nil {
		l.Debug("Unable to write to the file", esl.Error(errIO))
		return nil, es_filesystem_local.NewError(errIO)
	}

	if errIO := os.Chtimes(target.Path(), time.Now(), source.ModTime()); errIO != nil {
		l.Debug("Unable to modify time", esl.Error(err))
	}

	return z.target.Info(target)
}
