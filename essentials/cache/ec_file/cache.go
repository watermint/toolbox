package ec_file

import (
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"path/filepath"
)

// File caches remote file to local file system.
type File interface {
	// Get retrieve file from remote and cache to local file system.
	// namespace is the namespace of the cache.
	// name is the file name of the cache.
	// url is the remote file URL.
	Get(namespace, name, url string) (path string, err error)
}

func New(cacheRoot string, l esl.Logger) File {
	return &fileImpl{
		cacheRoot: cacheRoot,
		logger:    l.With(esl.String("cacheRoot", cacheRoot)),
	}
}

type fileImpl struct {
	cacheRoot string
	logger    esl.Logger
}

func (z fileImpl) Get(namespace, name, url string) (path string, err error) {
	l := z.logger.With(esl.String("namespace", namespace), esl.String("name", name), esl.String("url", url))
	cacheNamespacePath := filepath.Join(z.cacheRoot, namespace)
	cacheFilePath := filepath.Join(cacheNamespacePath, name)
	if _, err := os.Lstat(cacheFilePath); err == nil {
		l.Debug("Cache hit", esl.String("path", cacheFilePath))
		return cacheFilePath, nil
	}

	l.Debug("Cache miss", esl.String("path", cacheFilePath))
	if err := os.MkdirAll(cacheNamespacePath, 0755); err != nil {
		l.Debug("Unable to create cache directory", esl.Error(err))
		return "", err
	}

	if err := es_download.Download(l, url, cacheFilePath); err != nil {
		l.Debug("Unable to download", esl.Error(err))
		return "", err
	}

	return cacheFilePath, nil
}
