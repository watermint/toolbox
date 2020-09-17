package es_filesystem_model

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"path/filepath"
)

func NewPath(path string) es_filesystem.Path {
	return &Path{
		path: path,
	}
}

type Path struct {
	path string
}

func (z Path) IsRoot() bool {
	return z.path == "/"
}

func (z Path) Ancestor() es_filesystem.Path {
	return NewPath(filepath.ToSlash(filepath.Dir(z.path)))
}

func (z Path) AsData() es_filesystem.PathData {
	return es_filesystem.PathData{
		FileSystemType: FileSystemTypeModel,
		EntryPath:      z.Path(),
		EntryShard:     z.Shard().AsData(),
		Attributes:     map[string]interface{}{},
	}
}

func (z Path) Base() string {
	return filepath.Base(z.path)
}

func (z Path) Path() string {
	return z.path
}

func (z Path) Shard() es_filesystem.Shard {
	return es_filesystem.ShardData{
		ShardId: "",
	}
}

func (z Path) Descendant(pathFragment ...string) es_filesystem.Path {
	fragments := make([]string, 0)
	fragments = append(fragments, z.path)
	fragments = append(fragments, pathFragment...)
	return NewPath(filepath.Join(fragments...))
}

func (z Path) Rel(path es_filesystem.Path) (string, es_filesystem.FileSystemError) {
	rel, err := filepath.Rel(z.path, path.Path())
	if err != nil {
		return "", NewError(err, ErrorTypeOther)
	}
	return rel, nil
}
