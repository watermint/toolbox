package es_gzip

import (
	"compress/gzip"
	"github.com/watermint/toolbox/essentials/log/es_fallback"
	"go.uber.org/zap"
	"io"
	"os"
)

const (
	SuffixCompress = ".gz"
)

func Compress(path string) (err error) {
	cp := path + SuffixCompress
	l := es_fallback.Fallback().With(zap.String("path", cp))
	l.Debug("Compress")
	s, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open source", zap.Error(err))
		return err
	}
	c, err := os.Create(cp)
	if err != nil {
		l.Debug("Unable to create compressed", zap.Error(err))
		_ = s.Close()
		return err
	}

	zc := gzip.NewWriter(c)

	if _, err := io.Copy(zc, s); err != nil {
		l.Debug("Unable to copy", zap.Error(err))
		_ = c.Close()
		if err = os.Remove(cp); err != nil {
			l.Debug("unable to remove", zap.Error(err))
		}
		return err
	}
	if err := zc.Flush(); err != nil {
		l.Debug("Unable to flush", zap.Error(err))
		return err
	}

	// gzip writer's Close flushes & writes footer. Should catch an error.
	if err := zc.Close(); err != nil {
		l.Debug("Unable to close", zap.Error(err))
		return err
	}
	_ = s.Close()
	_ = c.Close()
	if err = os.Remove(path); err != nil {
		l.Debug("unable to remove", zap.Error(err))
	}
	return nil
}
