package es_filesystem_model

import (
	"errors"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"path/filepath"
)

const (
	FileSystemTypeModel = "model"
)

func NewFileSystem(root em_file.Node) es_filesystem.FileSystem {
	return &fileSystem{
		root: root,
	}
}

type fileSystem struct {
	root em_file.Node
}

func (z fileSystem) Path(data es_filesystem.PathData) (path es_filesystem.Path, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeModel {
		return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
	}

	return NewPath(data.EntryPath), nil
}

func (z fileSystem) Shard(data es_filesystem.ShardData) (namespace es_filesystem.Shard, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeModel {
		return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
	}

	return data, nil
}

func (z fileSystem) List(path es_filesystem.Path) (entries []es_filesystem.Entry, err es_filesystem.FileSystemError) {
	parent := em_file.ResolvePath(z.root, path.Path())
	if parent == nil {
		return nil, NewError(errors.New("not found"), ErrorTypePathNotFound)
	}
	switch n := parent.(type) {
	case em_file.Folder:
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
	node := em_file.ResolvePath(z.root, path.Path())
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
	parent := em_file.ResolvePath(z.root, parentPath)
	if parent == nil {
		return NewError(errors.New("not found"), ErrorTypePathNotFound)
	}
	switch n := parent.(type) {
	case em_file.Folder:
		found := n.Delete(filepath.Base(path.Base()))
		if !found {
			return NewError(errors.New("not found"), ErrorTypePathNotFound)
		}
		return nil
	}
	return NewError(errors.New("not found"), ErrorTypePathNotFound)
}

func (z fileSystem) CreateFolder(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	if em_file.CreateFolder(z.root, path.Path()) {
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
		contentSeed := es_number.New(data.Attributes[em_file.ExtraDataContentSeed]).Int64()
		entry = Entry{
			node:      em_file.NewFile(data.Name(), data.Size(), data.ModTime(), contentSeed),
			EntryData: data,
		}
		return

	case data.IsFolder():
		entry = Entry{
			node:      em_file.NewFolder(data.Name(), []em_file.Node{}),
			EntryData: data,
		}
		return
	}
	return nil, NewError(errors.New("invalid type"), ErrorTypeInvalidEntryDataFormat)
}

func (z fileSystem) FileSystemType() string {
	return FileSystemTypeModel
}
