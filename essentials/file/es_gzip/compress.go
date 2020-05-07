package es_gzip

import (
	"compress/gzip"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"io"
	"os"
)

const (
	SuffixCompress = ".gz"
)

func Compress(path string) (err error) {
	cp := path + SuffixCompress
	l := es_log.Default().With(es_log.String("path", cp))
	l.Debug("Compress")
	s, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open source", es_log.Error(err))
		return err
	}
	c, err := os.Create(cp)
	if err != nil {
		l.Debug("Unable to create compressed", es_log.Error(err))
		_ = s.Close()
		return err
	}

	zc := gzip.NewWriter(c)

	if _, err := io.Copy(zc, s); err != nil {
		l.Debug("Unable to copy", es_log.Error(err))
		_ = c.Close()
		if err = os.Remove(cp); err != nil {
			l.Debug("unable to remove", es_log.Error(err))
		}
		return err
	}
	if err := zc.Flush(); err != nil {
		l.Debug("Unable to flush", es_log.Error(err))
		return err
	}

	// gzip writer's Close flushes & writes footer. Should catch an error.
	if err := zc.Close(); err != nil {
		l.Debug("Unable to close", es_log.Error(err))
		return err
	}
	_ = s.Close()
	_ = c.Close()
	if err = os.Remove(path); err != nil {
		l.Debug("unable to remove", es_log.Error(err))
	}
	return nil
}
