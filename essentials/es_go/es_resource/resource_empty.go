package es_resource

import (
	"github.com/watermint/toolbox/essentials/http/es_filesystem"
	"net/http"
)

func EmptyResource() Resource {
	return &resEmpty{}
}

// empty resource
type resEmpty struct {
}

func (z resEmpty) Bytes(key string) (bin []byte, err error) {
	return nil, ErrorAlwaysFail
}

func (z resEmpty) HttpFileSystem() http.FileSystem {
	return es_filesystem.Empty{}
}
