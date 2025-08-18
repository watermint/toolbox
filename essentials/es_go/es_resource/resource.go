package es_resource

import (
	"errors"
	"net/http"
)

const (
	BundleResource = "resources"
	BundleWeb      = "web"
)

type Resource interface {
	Bytes(key string) (bin []byte, err error)
	HttpFileSystem() http.FileSystem
}

var (
	ErrorAlwaysFail = errors.New("always fail")
	ErrorNotFound   = errors.New("not found")
)
