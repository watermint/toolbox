package es_filesystem_connector

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_tree"
)

func NewModelToModel(l esl.Logger, sourceRoot, targetRoot em_tree.Folder) es_filesystem.Connector {
	return &modelToModelConn{
		l:          l,
		sourceRoot: sourceRoot,
		targetRoot: targetRoot,
	}
}

type modelToModelConn struct {
	l          esl.Logger
	sourceRoot em_tree.Folder
	targetRoot em_tree.Folder
}

func (z modelToModelConn) Copy(source es_filesystem.Entry, target es_filesystem.Path) (err es_filesystem.FileSystemError) {
	l := z.l.With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")

	sourceNode := em_tree.ResolvePath(z.sourceRoot, source.Path().Path())
	if sourceNode == nil {
		l.Debug("Unable to find the source node")
		return es_filesystem_model.NewError(errors.New("source node not found"), es_filesystem_model.ErrorTypePathNotFound)
	}

	if sourceNode.Type() != em_tree.FileNode {
		l.Debug("Node is not a file")
		return es_filesystem_model.NewError(errors.New("source node is not a file"), es_filesystem_model.ErrorTypeOther)
	}

	targetFolderPath := target.Ancestor()
	targetFolder := em_tree.ResolvePath(z.targetRoot, targetFolderPath.Path())
	if targetFolder == nil {
		if !em_tree.CreateFolder(z.targetRoot, targetFolderPath.Path()) {
			l.Debug("Unable to create folder", esl.Error(err))
			return es_filesystem_model.NewError(err, es_filesystem_model.ErrorTypeOther)
		}
		targetFolder = em_tree.ResolvePath(z.targetRoot, targetFolderPath.Path())
		if targetFolder == nil {
			l.Debug("Unable to resolve target path")
			return es_filesystem_model.NewError(errors.New("unable to resolve target path"), es_filesystem_model.ErrorTypePathNotFound)
		}
	}

	if targetFolder.Type() != em_tree.FolderNode {
		l.Debug("target nod is not a folder")
		return es_filesystem_model.NewError(errors.New("target node is not a folder"), es_filesystem_model.ErrorTypeConflict)
	}

	targetFolder.(em_tree.Folder).Add(sourceNode)

	return nil
}
