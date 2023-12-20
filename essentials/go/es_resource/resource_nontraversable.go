package es_resource

import (
	"embed"
	"github.com/watermint/toolbox/essentials/http/es_filesystem"
	"net/http"
	"path/filepath"
)

func NewNonTraversableResource(prefix string, fs embed.FS) Resource {
	return &resNonTraversableBox{
		prefix: prefix,
		fs:     fs,
	}
}

// embed.FS wrapper, but do not return http.FileSystem
type resNonTraversableBox struct {
	prefix string
	fs     embed.FS
}

func (z resNonTraversableBox) Bytes(key string) (bin []byte, err error) {
	return z.fs.ReadFile(filepath.ToSlash(filepath.Join(z.prefix, key)))
}

func (z resNonTraversableBox) HttpFileSystem() http.FileSystem {
	// Always return empty to prevent resource directory listing
	return es_filesystem.Empty{}
}
