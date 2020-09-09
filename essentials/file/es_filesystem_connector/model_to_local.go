package es_filesystem_connector

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
	return &modelToLocalConn{
		l:          l,
		sourceRoot: sourceRoot,
	}
}

type modelToLocalConn struct {
	l          esl.Logger
	sourceRoot em_tree.Folder
}

func (z modelToLocalConn) Copy(source es_filesystem.Entry, target es_filesystem.Path) (err es_filesystem.FileSystemError) {
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

	content := sourceNode.(em_tree.File).Content()

	if errIO := ioutil.WriteFile(target.Path(), content, 0644); errIO != nil {
		l.Debug("Unable to write to the file", esl.Error(errIO))
		return es_filesystem_local.NewError(errIO)
	}

	if errIO := os.Chtimes(target.Path(), time.Now(), source.ModTime()); errIO != nil {
		l.Debug("Unable to modify time", esl.Error(err))
	}

	return nil
}
