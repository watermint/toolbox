package es_resource

import (
	"net/http"
)

func NewMergedResource(resources ...Resource) Resource {
	return &mergedResource{
		resources: resources,
	}
}

type mergedResource struct {
	resources []Resource
}

func (z mergedResource) Bytes(key string) (bin []byte, err error) {
	for _, r := range z.resources {
		bin, err = r.Bytes(key)
		if err == nil {
			return bin, nil
		}
	}
	return nil, ErrorNotFound
}

func (z mergedResource) HttpFileSystem() http.FileSystem {
	fs := make([]http.FileSystem, 0)
	for _, r := range z.resources {
		fs = append(fs, r.HttpFileSystem())
	}
	return mergedHttpFileSystem{
		fileSystems: fs,
	}
}

type mergedHttpFileSystem struct {
	fileSystems []http.FileSystem
}

func (z mergedHttpFileSystem) Open(name string) (http.File, error) {
	for _, fs := range z.fileSystems {
		f, err := fs.Open(name)
		if err == nil {
			return f, nil
		}
	}
	return nil, ErrorNotFound
}
