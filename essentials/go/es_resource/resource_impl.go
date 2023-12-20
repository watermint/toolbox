package es_resource

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"
)

func NewResource(prefix string, fs embed.FS) Resource {
	return &resBox{
		prefix: prefix,
		fs:     fs,
	}
}

// go embed wrapper
type resBox struct {
	prefix string
	fs     embed.FS
}

func (z resBox) Bytes(key string) (bin []byte, err error) {
	return z.fs.ReadFile(filepath.ToSlash(filepath.Join(z.prefix, key)))
}

func (z resBox) HttpFileSystem() http.FileSystem {
	f, err := fs.Sub(z.fs, z.prefix)
	if err != nil {
		panic(err)
	}
	return http.FS(f)
}
