package es_filesystem_model

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"io/ioutil"
	"time"
)

func NewEntry(path string, node em_file.Node) es_filesystem.Entry {
	switch n := node.(type) {
	case em_file.Folder:
		return &Entry{
			node: node,
			EntryData: es_filesystem.EntryData{
				FileSystemType: FileSystemTypeModel,
				EntryName:      n.Name(),
				EntryPath:      path,
				EntrySize:      0,
				EntryModTime:   time.Time{},
				EntryIsFile:    false,
				EntryIsFolder:  true,
				Attributes:     node.ExtraData(),
			},
		}

	case em_file.File:
		return &Entry{
			node: node,
			EntryData: es_filesystem.EntryData{
				FileSystemType: FileSystemTypeModel,
				EntryName:      n.Name(),
				EntryPath:      path,
				EntrySize:      n.Size(),
				EntryModTime:   n.ModTime(),
				EntryIsFile:    true,
				EntryIsFolder:  false,
				Attributes:     node.ExtraData(),
			},
		}
	}
	panic("unsupported node type")
}

type Entry struct {
	node em_file.Node
	es_filesystem.EntryData
}

func (z Entry) AsData() es_filesystem.EntryData {
	z.EntryData.Attributes = z.node.ExtraData()
	return z.EntryData
}

func (z Entry) Path() es_filesystem.Path {
	return NewPath(z.EntryPath)
}

func (z Entry) ContentHash() (string, es_filesystem.FileSystemError) {
	switch n := z.node.(type) {
	case em_file.File:
		content := n.Content()
		hash, err := dbx_util.ContentHash(ioutil.NopCloser(bytes.NewReader(content)), int64(len(content)))
		if err != nil {
			return "", NewError(errors.New("invalid type"), ErrorTypeOther)
		}
		return hash, nil

	default:
		return "", NewError(errors.New("invalid type"), ErrorTypeOther)
	}
}
