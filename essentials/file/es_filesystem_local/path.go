package es_filesystem_local

import (
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"path/filepath"
)

func NewPath(path string) es_filesystem.Path {
	var p string
	if path == "" {
		p = "/"
	} else {
		var err error
		p, err = filepath.Abs(path)
		if err != nil {
			p = path
		}
	}
	return &fsPath{
		path: filepath.ToSlash(filepath.Clean(p)),
	}
}

type fsPath struct {
	path string
}

func (z fsPath) IsRoot() bool {
	return z.path == "/" || z.path == ""
}

func (z fsPath) Ancestor() es_filesystem.Path {
	return NewPath(filepath.Dir(z.path))
}

func (z fsPath) AsData() es_filesystem.PathData {
	return es_filesystem.PathData{
		FileSystemType: FileSystemTypeLocal,
		EntryPath:      z.path,
		EntryNamespace: z.Namespace().AsData(),
		Attributes:     map[string]interface{}{},
	}
}

func (z fsPath) Base() string {
	return filepath.Base(z.path)
}

func (z fsPath) Path() string {
	return z.path
}

func (z fsPath) Namespace() es_filesystem.Namespace {
	return es_filesystem.NamespaceData{
		NamespaceId: filepath.VolumeName(z.path),
	}
}

func (z fsPath) Descendant(pathFragment ...string) es_filesystem.Path {
	fragments := make([]string, 0)
	fragments = append(fragments, z.path)
	fragments = append(fragments, pathFragment...)
	return NewPath(filepath.Join(fragments...))
}

func (z fsPath) Rel(path es_filesystem.Path) (string, es_filesystem.FileSystemError) {
	r, err := es_filepath.Rel(z.path, path.Path())
	return r, NewError(err)
}
