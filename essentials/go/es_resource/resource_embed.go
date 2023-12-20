package es_resource

import (
	"github.com/watermint/toolbox/essentials/http/es_filesystem"
	"net/http"
)

func NewEmbedResource(resources map[string][]byte) Resource {
	return &embedResource{
		resources: resources,
	}
}

type embedResource struct {
	resources map[string][]byte
}

func (z embedResource) Bytes(key string) (bin []byte, err error) {
	if b, ok := z.resources[key]; ok {
		return b, nil
	}
	return nil, ErrorNotFound
}

func (z embedResource) HttpFileSystem() http.FileSystem {
	return es_filesystem.Empty{}
}
