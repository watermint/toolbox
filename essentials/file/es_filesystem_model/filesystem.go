package es_filesystem_model

import (
	"errors"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/model/em_tree"
	"path/filepath"
)

const (
	FileSystemTypeModel = "model"
)

func NewFileSystem(root em_tree.Node) es_filesystem.FileSystem {
	return &fileSystem{
		root: root,
	}
}

type fileSystem struct {
	root em_tree.Node
}

func (z fileSystem) Path(data es_filesystem.PathData) (path es_filesystem.Path, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeModel {
		return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
	}

	return NewPath(data.EntryPath), nil
}

func (z fileSystem) Namespace(data es_filesystem.NamespaceData) (namespace es_filesystem.Namespace, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeModel {
		return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
	}

	return data, nil
}

func (z fileSystem) List(path es_filesystem.Path) (entries []es_filesystem.Entry, err es_filesystem.FileSystemError) {
	parent := em_tree.ResolvePath(z.root, path.Path())
	if parent == nil {
		return nil, NewError(errors.New("not found"), ErrorTypePathNotFound)
	}
	switch n := parent.(type) {
	case em_tree.Folder:
		nodeEntries := n.Descendants()
		entries = make([]es_filesystem.Entry, 0)
		for _, nodeEntry := range nodeEntries {
			entries = append(entries, NewEntry(path.Descendant(nodeEntry.Name()).Path(), nodeEntry))
		}
		return entries, nil
	}
	return nil, NewError(errors.New("not found"), ErrorTypePathNotFound)
}

func (z fileSystem) Info(path es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	node := em_tree.ResolvePath(z.root, path.Path())
	if node == nil {
		return nil, NewError(errors.New("not found"), ErrorTypePathNotFound)
	}
	return NewEntry(path.Path(), node), nil
}

func (z fileSystem) Delete(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	if path.IsRoot() {
		return NewError(errors.New("unable to delete root"), ErrorTypeOther)
	}
	parentPath := path.Ancestor().Path()
	parent := em_tree.ResolvePath(z.root, parentPath)
	if parent == nil {
		return NewError(errors.New("not found"), ErrorTypePathNotFound)
	}
	switch n := parent.(type) {
	case em_tree.Folder:
		found := n.Delete(filepath.Base(path.Base()))
		if !found {
			return NewError(errors.New("not found"), ErrorTypePathNotFound)
		}
		return nil
	}
	return NewError(errors.New("not found"), ErrorTypePathNotFound)
}

func (z fileSystem) CreateFolder(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	if em_tree.CreateFolder(z.root, path.Path()) {
		return nil
	} else {
		return NewError(errors.New("conflict"), ErrorTypeConflict)
	}
}

func (z fileSystem) Entry(data es_filesystem.EntryData) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeModel {
		return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
	}

	switch {
	case data.IsFile():
		contentSeed := es_number.New(data.Attributes[em_tree.ExtraDataContentSeed]).Int64()
		entry = Entry{
			node:      em_tree.NewFile(data.Name(), data.Size(), data.ModTime(), contentSeed),
			EntryData: data,
		}
		return

	case data.IsFolder():
		entry = Entry{
			node:      em_tree.NewFolder(data.Name(), []em_tree.Node{}),
			EntryData: data,
		}
		return
	}
	return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
}

func (z fileSystem) FileSystemType() string {
	return FileSystemTypeModel
}
